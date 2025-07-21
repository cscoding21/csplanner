package comment

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/providers/postgres"
	"csserver/internal/utils/quilljs"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ---This is the name of the object in the database
const ReactionIdentifier = postgres.TableName("reaction")

const (
	Likes       CommentReactionType = "likes"
	Dislikes    CommentReactionType = "dislikes"
	Loves       CommentReactionType = "loves"
	LaughsAt    CommentReactionType = "laughs_at"
	Acknowledge CommentReactionType = "acknowledges"
)

// AddComment adds a comment to a project.  Is a wrapper for CreateComment
func (s *CommentService) AddComment(ctx context.Context, comment Comment) (common.UpdateResult[*common.BaseModel[Comment]], error) {
	c, err := s.CreateComment(ctx, comment)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Comment]](nil, err)
	}

	cc := *c.Object
	err = s.pubsub.StreamPublish(ctx,
		string(CommentIdentifier),
		"comment", "created",
		map[string]any{
			"text":       cc.Data.Text,
			"id":         cc.ID,
			"project_id": cc.Data.ProjectID,
		})
	if err != nil {
		log.Errorf("StreamPublish error: %s", err)
	} else {
		log.Warnf("StreamPublish success!")
	}

	return c, nil
}

// AddActivityComment add a comment for an activity
func (s *CommentService) AddActivityComment(ctx context.Context, comment Comment) (common.UpdateResult[*common.BaseModel[Comment]], error) {
	comment.IsActivity = true
	return s.CreateComment(ctx, comment)
}

// AddComment create a new comment for a project
func (s *CommentService) AddCommentReply(ctx context.Context, comment Comment, parentCommentID string) (common.UpdateResult[*common.BaseModel[Comment]], error) {
	log.Infof("RETRIEVING PARENT COMMENT: %v", parentCommentID)
	parentComment, err := s.GetCommentByID(ctx, parentCommentID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Comment]](nil, err)
	}

	fmt.Println(parentComment)
	log.Debugf("CREATE COMMENT REPLY: %v", comment)
	comment.ProjectID = parentComment.Data.ProjectID

	//---clear out the project id for replies
	outComment, err := postgres.UpdateObjectWithParent(ctx, s.db, comment, CommentIdentifier, "", &parentCommentID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Comment]](nil, err)
	}

	replyDelta, err := quilljs.FromJSON([]byte(comment.Text))
	if err != nil {
		log.Errorf("error creating reply delta: %v", err)
	}

	fmt.Println("replyDelta", replyDelta)

	cc := *outComment
	err = s.pubsub.StreamPublish(ctx,
		string(CommentIdentifier),
		"reply", "created",
		map[string]any{
			"text":           cc.Data.Text,
			"id":             cc.ID,
			"project_id":     parentComment.Data.ProjectID,
			"parent_user_id": parentComment.CreatedBy,
		})
	if err != nil {
		log.Errorf("StreamPublish error: %s", err)
	} else {
		log.Warnf("StreamPublish success!")
	}

	return common.NewSuccessUpdateResult(outComment)
}

// ModifyComment update the text of an existing comment
func (s *CommentService) ModifyComment(ctx context.Context, comment Comment) (common.UpdateResult[*common.BaseModel[Comment]], error) {
	userEmail := config.GetUserEmailFromContext(ctx)

	existingComment, err := s.GetCommentByID(ctx, comment.ID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Comment]](nil, err)
	}

	if existingComment.Data.IsActivity {
		err := fmt.Errorf("activity comments cannot be modified")
		return common.NewFailingUpdateResult[*common.BaseModel[Comment]](nil, err)
	}

	if userEmail != existingComment.CreatedBy {
		err := fmt.Errorf("user %v is not allowed to update comment %v", userEmail, existingComment.ID)
		return common.NewFailingUpdateResult[*common.BaseModel[Comment]](nil, err)
	}

	existingComment.Data.Text = comment.Text
	existingComment.Data.IsEdited = true

	result, err := postgres.UpdateObjectWithParent(ctx, s.db, existingComment.Data, CommentIdentifier, existingComment.ID, existingComment.ParentID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Comment]](nil, err)
	}

	s.pubsub.StreamPublish(ctx, string(CommentIdentifier), "comment", "updated", result)

	return common.NewSuccessUpdateResult(result)
}

// FindProjectComments finds all activity for a given project
func (s *CommentService) FindProjectComments(ctx context.Context, projectID string) (common.PagedResults[common.BaseModel[Comment]], error) {
	pagedResults := common.NewPagedResultsForAllRecords[common.BaseModel[Comment]]()
	filters := common.Filters{}
	filters.AddFilter(common.Filter{
		Key:       "data.project_id",
		Value:     projectID,
		Operation: common.FilterOperationEqual,
	})

	sql := `select * 
			from comment 
			where true 
				and deleted_at is null 
				and data->>'project_id' = $1
			order by updated_at`

	results, err := postgres.FindPagedObjects[Comment](ctx, s.db, sql, pagedResults.Pagination, filters, filters.GetFiltersOrderedValues())
	if err != nil {
		return pagedResults, err
	}

	reactions, err := s.FindProjectReactions(ctx, projectID)
	if err != nil {
		return pagedResults, err
	}
	resultMap := commentsToMap(results.Results)
	applyReactionsToCommentList(resultMap, reactions.Results)

	graphedResults, err := commentResultsToGraph(resultMap)
	if err != nil {
		return pagedResults, err
	}

	results.Results = graphedResults

	return results, nil
}

func commentsToMap(results []common.BaseModel[Comment]) map[string]common.BaseModel[Comment] {
	out := make(map[string]common.BaseModel[Comment])

	for _, r := range results {
		out[r.ID] = r
	}

	return out
}

func commentResultsToGraph(results map[string]common.BaseModel[Comment]) ([]common.BaseModel[Comment], error) {
	out := []common.BaseModel[Comment]{}
	pm := make(map[string]common.BaseModel[Comment])

	for _, r := range results {
		if r.ParentID == nil {
			pm[r.ID] = r
		}
	}

	for _, r := range results {
		if r.ParentID != nil {
			pid := *r.ParentID
			item, ok := pm[pid]
			if !ok {
				//---should be impossible
				log.Warnf("parent comment not found for id %s when building graph", pid)
			}

			item.Data.CommentReplies = append(item.Data.CommentReplies, r)
			pm[pid] = item
		}
	}

	for _, v := range pm {
		out = append(out, v)
	}

	return out, nil
}

func applyReactionsToCommentList(resultMap map[string]common.BaseModel[Comment], reactions []common.BaseModel[Reaction]) {
	for _, react := range reactions {
		item, ok := resultMap[*react.ParentID]
		if ok {
			switch react.Data.Type {
			case (Likes):
				item.Data.Likes = append(item.Data.Likes, react.CreatedBy)
			case (Dislikes):
				item.Data.Dislikes = append(item.Data.Dislikes, react.CreatedBy)
			case (Loves):
				item.Data.Loves = append(item.Data.Loves, react.CreatedBy)
			case (Acknowledge):
				item.Data.Acknowledges = append(item.Data.Acknowledges, react.CreatedBy)
			case (LaughsAt):
				item.Data.LaughsAt = append(item.Data.LaughsAt, react.CreatedBy)
			}

			resultMap[*react.ParentID] = item
		}
	}
}

// GetCommentThread return the entire comment thread for an id whether it's a top-level or child
func (s *CommentService) GetCommentThread(ctx context.Context, commentID string) (*common.BaseModel[Comment], error) {
	comment, err := s.GetCommentByID(ctx, commentID)
	if err != nil {
		return nil, err
	}

	if comment.ParentID != nil {
		//---this is a child...we want the parent
		comment, err = s.GetCommentByID(ctx, *comment.ParentID)
		if err != nil {
			return nil, err
		}
	}

	sql := fmt.Sprintf(`select * from %s c where c.id = $1 or c.parent_id = $1 order by updated_at`, CommentIdentifier)
	pa := common.NewPagedResultsForAllRecords[Comment]()
	pa.Filters.AddFilter(common.Filter{Key: "id", Value: comment.ID, Operation: common.FilterOperationEqual})

	threadComments, err := postgres.FindPagedObjects[Comment](ctx, s.db, sql, pa.Pagination, pa.Filters, pa.Filters.GetFiltersOrderedValues())
	if err != nil {
		return nil, err
	}

	reactions, err := s.FindProjectReactions(ctx, comment.Data.ProjectID)
	if err != nil {
		return nil, err
	}
	resultMap := commentsToMap(threadComments.Results)
	applyReactionsToCommentList(resultMap, reactions.Results)

	graphedResults, err := commentResultsToGraph(resultMap)
	if err != nil {
		return nil, err
	}

	return &graphedResults[0], nil
}

// FindProjectComments finds all activity for a given project
func (s *CommentService) FindCommentReplies(ctx context.Context, commentID string) (common.PagedResults[common.BaseModel[Comment]], error) {
	return postgres.FindAllChildren[Comment](ctx, s.db, CommentIdentifier, commentID)
}

// ToggleCommentReaction if a user has emoted to the comment...remove the existing emote.  Otherwise add it
func (s *CommentService) ToggleCommentReaction(ctx context.Context, projectID string, parentID *string, commentID string, reaction CommentReactionType) error {
	userEmail := config.GetUserEmailFromContext(ctx)
	sql := fmt.Sprintf(`select * from %s where created_by = $1 and parent_id = $2 and data->>'type' = $3`, ReactionIdentifier)
	params := []any{userEmail, commentID, reaction}

	log.Warnf("Existing reaction SQL: %s", sql)
	log.Warnf("Existing params: %v", params)

	existingReaction, err := postgres.GetObject[common.BaseModel[Reaction]](ctx, s.db, ReactionIdentifier, sql, params...)
	if err != nil {
		return err
	}

	log.Warnf("Existing reaction: %v", existingReaction)

	if existingReaction != nil {
		if !strings.EqualFold(existingReaction.CreatedBy, userEmail) {
			return fmt.Errorf("user %s cannot change the reaction of user %s", userEmail, existingReaction.CreatedBy)
		}

		err = s.deleteCommentReaction(ctx, existingReaction.ID)
		if err != nil {
			return err
		}

		return nil
	}

	err = s.createCommentReaction(ctx, projectID, parentID, commentID, reaction)
	if err != nil {
		return err
	}

	return nil
}

// CreateCommentEmote add an emote to an comment
func (s *CommentService) createCommentReaction(ctx context.Context, projectID string, parentID *string, commentID string, reaction CommentReactionType) error {
	obj := Reaction{Type: reaction, ProjectID: projectID}

	log.Warnf("New reaction: %v", obj)

	_, err := postgres.UpdateObjectWithParent(ctx, s.db, obj, ReactionIdentifier, "", &commentID)
	return err
}

// DeleteCommentEmote deletes an emote previously creatd by a user
func (s *CommentService) deleteCommentReaction(ctx context.Context, id string) error {
	return postgres.Delete(ctx, s.db, ReactionIdentifier, id)
}

// RemoveComment soft-delete a comment based on its ID
func (s *CommentService) RemoveComment(ctx context.Context, commentID string) error {
	existingComment, err := s.GetCommentByID(ctx, commentID)
	if err != nil {
		return err
	}

	if existingComment.Data.IsActivity {
		return fmt.Errorf("activity comments cannot be deleted")
	}

	return postgres.SoftDelete(ctx, s.db, ReactionIdentifier, commentID)
}

// FindProjectReactions find all reactions related to project comments
func (s *CommentService) FindProjectReactions(ctx context.Context, projectID string) (common.PagedResults[common.BaseModel[Reaction]], error) {
	sql := fmt.Sprintf(`select * from %s where data->>'project_id' = $1`, ReactionIdentifier)
	pf := common.NewPagedResultsForAllRecords[common.BaseModel[Reaction]]()
	pf.Filters.AddFilter(common.Filter{Key: "data.project_id", Value: projectID, Operation: common.FilterOperationEqual})

	return postgres.FindPagedObjects[Reaction](ctx, s.db, sql, pf.Pagination, pf.Filters, pf.Filters.GetFiltersOrderedValues())
}

/*
// extractMentions extracts all mentions from a delta
func extractMentions(delta quilljs.Delta, commentID, projectID, contextUserID, contextUserName string) []Mention {
	mentions := []Mention{}

	for _, o := range delta.Ops {
		if o.InsertEmbed != nil {
			if o.InsertEmbed.Key == "mention" {
				m := Mention{
					UserEmail:       o.InsertEmbed.Value.(map[string]interface{})["id"].(string),
					UserName:        o.InsertEmbed.Value.(map[string]interface{})["value"].(string),
					CommentID:       commentID,
					ProjectID:       projectID,
					MentionedByID:   contextUserID,
					MentionedByName: contextUserName,
					Text:            quilljs.DeltaToString(delta),
				}

				mentions = append(mentions, m)
			}
		}
	}

	return mentions
}
*/
