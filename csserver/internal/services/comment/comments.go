package comment

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/marshal"
	"csserver/internal/utils/quilljs"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

type CommentRelationshipType string

const (
	Likes       CommentRelationshipType = "likes"
	Dislikes    CommentRelationshipType = "dislikes"
	Loves       CommentRelationshipType = "loves"
	LaughsAt    CommentRelationshipType = "laughs_at"
	Acknowledge CommentRelationshipType = "acknowledges"

	Mentioned CommentRelationshipType = "mentioned"

	IsAReplyTo CommentRelationshipType = "is_a_reply_to"
	BelongsTo  CommentRelationshipType = "belongsto"
)

// AddComment adds a comment to a project.  Is a wrapper for CreateComment
func (s *CommentService) AddComment(ctx context.Context, comment Comment) (*common.UpdateResult[Comment], error) {
	c, err := s.CreateComment(ctx, &comment)
	if err != nil {
		return nil, err
	}

	//handleMentions(ctx, c, resource, c.ID)

	//---create the realtionship for the user that created the project
	//activityDelta := getAddCommentActivityDelta(*user, project)

	// _, err = activityService.CreateActivity(ctx, c.ID, activityDelta, utils.GetTruncatedText(comment.Text, 100), *project.ID)
	// if err != nil {
	// 	log.Error(err)
	// 	return nil, err
	// }
	if comment.ProjectID != "" {
		_, err = s.DBClient.RelateTo(ctx, c.Object.ID, comment.ProjectID, string(BelongsTo))
		if err != nil {
			return nil, err
		}
	}

	return &c, nil
}

// AddComment create a new comment for a project
func (s *CommentService) AddCommentReply(ctx context.Context, comment Comment, parentCommentID string) (*common.UpdateResult[Comment], error) {
	//userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	// user, err := userService.GetCurrentUser(ctx)
	// if err != nil {
	// 	log.Error(err)
	// 	return nil, err
	// }
	// resource, err := projectService.GetResource(ctx, user.ID)
	// if err != nil {
	// 	log.Error(err)
	// 	return nil, err
	// }

	log.Infof("RETRIEVING PARENT COMMENT: %v", parentCommentID)
	parentComment, err := s.GetCommentByID(ctx, parentCommentID)
	if err != nil {
		return nil, err
	}

	// parentCommentUser, err := projectService.GetResource(ctx, parentComment.ControlFields.CreatedBy)
	// if err != nil {
	// 	return nil, err
	// }

	// project, err := database.GetObjectById[projectModels.Project](ctx, parentComment.ProjectID)
	// if err != nil {
	// 	return nil, err
	// }

	log.Debugf("CREATE COMMENT REPLY: %v", comment)

	//---clear out the project id for replies
	comment.ProjectID = ""
	c, err := s.CreateComment(ctx, &comment)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// parentCommentIDScalar, err := database.GetScalar("select out from belongsto where in = $parentCommentID", "out", map[string]interface{}{"parentCommentID": c.ID})
	// if err != nil {
	// 	log.Error(err)
	// 	return nil, err
	// }
	// if len(fmt.Sprint(parentCommentIDScalar)) > 0 {
	// 	parentCommentID = parentCommentIDScalar.(string)
	// 	log.Infof("UPDATING PARENT COMMENT ID TO %s", parentCommentID)
	// }

	// handleMentions(ctx, c, resource, parentCommentID)
	// pubsub.Publish(ctx, "notification", "comment", "replied", comment)

	replyDelta, err := quilljs.FromJSON([]byte(comment.Text))
	if err != nil {
		log.Errorf("error creating reply delta: %v", err)
	}

	fmt.Println("replyDelta", replyDelta)

	// notificationService.DispatchReplies(ctx, notificationModels.Notification{
	// 	UserID:                *parentCommentUser.UserID,
	// 	UserName:              parentCommentUser.Name,
	// 	Type:                  notificationType.Reply,
	// 	ContextID:             parentCommentID,
	// 	InitiatorName:         resource.Name,
	// 	InitiatorID:           *resource.ID,
	// 	InitiatorProfileImage: resource.ProfileImage,
	// 	Text:                  quilljs.DeltaToString(*replyDelta),
	// 	RecipientIsBot:        parentCommentUser.IsBot,
	// })

	//---create the realtionship for the user that created the project
	// activityDelta := getAddCommentReplyActivityDelta(*user, project)

	// _, err = activityService.CreateActivity(ctx, c.ID, activityDelta, utils.GetTruncatedText(comment.Text, 100), *project.ID)
	// if err != nil {
	// 	return nil, err
	// }

	_, err = s.DBClient.RelateTo(ctx, c.Object.ID, parentComment.ID, string(IsAReplyTo))
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// ModifyComment update the text of an existing comment
func (s *CommentService) ModifyComment(ctx context.Context, comment Comment) (*common.UpdateResult[Comment], error) {
	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	existingComment, err := s.GetCommentByID(ctx, comment.ID)
	if err != nil {
		return nil, err
	}

	if userEmail != existingComment.ControlFields.CreatedBy {
		err := fmt.Errorf("user %v is not allowed to update comment %v", userEmail, existingComment.ID)
		return nil, err
	}

	existingComment.Text = comment.Text
	existingComment.IsEdited = true

	c, err := s.UpdateComment(ctx, existingComment)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// FindProjectComments finds all activity for a given project
func (s *CommentService) FindProjectComments(ctx context.Context, projectID string) (*common.PagedResults[Comment], error) {
	pagingResults := common.NewPagedResultsForAllRecords[Comment]()
	sql := `SELECT *,
		<-likes<-user.id as likes,
		<-loves<-user.id as loves,
		<-dislikes<-user.id as dislikes,
		<-laughs_at<-user.id as laughs_at,
		<-acknowledges<-user.id as acknowledges
	FROM comment
	WHERE ->belongsto->(project where id = $project)
	ORDER BY control_fields.created_at
	`

	paging := common.NewPagedResultsForAllRecords[Comment]()
	filters := common.Filters{}
	filters.AddFilter(common.Filter{
		Key:       "project",
		Value:     projectID,
		Operation: common.FilterOperationEqual,
	})

	results, resultCount, err := s.DBClient.FindPagedObjects(sql, paging.Pagination, filters)
	if err != nil {
		return nil, err
	}
	pagingResults.Pagination.TotalResults = &resultCount

	unpacked, err := marshal.SurrealSmartUnmarshal[[]Comment](results)
	if err != nil {
		return nil, err
	}

	op := []Comment{}
	for _, c := range *unpacked {
		replies, err := s.FindCommentReplies(ctx, c.ID)
		if err != nil {
			break
		}

		if replies != nil && len(replies.Results) > 0 {
			c.Replies = append(c.Replies, replies.Results...)
		}

		op = append(op, c)
	}

	pagingResults.Results = op

	return &pagingResults, nil
}

// GetCommentThread return the entire comment thread for an id whether it's a top-level or child
func (s *CommentService) GetCommentThread(ctx context.Context, commentID string) (*Comment, error) {
	var outComment *Comment
	comment, err := s.GetCommentByID(ctx, commentID)
	if err != nil {
		return nil, err
	}

	filters := common.Filters{}
	filters.AddFilter(common.Filter{
		Key:       "commentID",
		Value:     commentID,
		Operation: common.FilterOperationEqual,
	})

	//---if this is not a top-level comment, then we need to get the parent
	if comment.ProjectID == "" {
		sql := `SELECT *
			FROM comment
			WHERE <-is_a_reply_to<-(comment where id = $commentID)
			`
		outCommentRaw, err := s.DBClient.GetObject(sql, filters.GetFiltersAsMap())
		if err != nil {
			return nil, err
		}

		log.Warnf("outCommentRaw: %v", outCommentRaw)
		outCommentSlice, err := marshal.SurrealSmartUnmarshal[[]Comment](outCommentRaw)
		if err != nil {
			return nil, err
		}

		log.Warnf("outCommentSlice: %v", outCommentSlice)
		if len(*outCommentSlice) > 0 {
			outComment = &(*outCommentSlice)[0]
		} else {
			return nil, fmt.Errorf("no comment thread found for comment %s", commentID)
		}
	} else {
		outComment = comment
	}

	replies, err := s.FindCommentReplies(ctx, outComment.ID)
	if err != nil {
		return nil, err
	}

	if replies != nil && replies.Results != nil && len(replies.Results) > 0 {
		outComment.Replies = replies.Results
	}

	return outComment, nil
}

// FindProjectComments finds all activity for a given project
func (s *CommentService) FindCommentReplies(ctx context.Context, commentID string) (*common.PagedResults[Comment], error) {
	pagingResults := common.NewPagedResultsForAllRecords[Comment]()
	sql := `SELECT *, 
		<-likes<-user.id as likes,
		<-loves<-user.id as loves,
		<-dislikes<-user.id as dislikes,
		<-laughs_at<-user.id as laughs_at,
		<-acknowledges<-user.id as acknowledges
	FROM comment 
	WHERE ->is_a_reply_to->(comment where id = $comment) 
	ORDER BY control_fields.created_at
	`

	paging := common.NewPagedResultsForAllRecords[Comment]()
	filters := common.Filters{}
	filters.AddFilter(common.Filter{
		Key:       "comment",
		Value:     commentID,
		Operation: common.FilterOperationEqual,
	})

	results, _, err := s.DBClient.FindPagedObjects(sql, paging.Pagination, filters)
	if err != nil {
		return nil, err
	}

	unpacked, err := marshal.SurrealSmartUnmarshal[[]Comment](results)
	if err != nil {
		return nil, err
	}

	pagingResults.Results = *unpacked

	return &pagingResults, nil
}

// ToggleCommentEmote if a user has emoted to the comment...remove the existing emote.  Otherwise add it
func (s *CommentService) ToggleCommentEmote(ctx context.Context, commentID string, emote CommentRelationshipType) error {
	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	existingRelation, err := s.DBClient.TestRelationshipExist(ctx, userEmail, commentID, string(emote))
	if err != nil {
		return err
	}

	if existingRelation != nil {
		err = s.DeleteCommentEmote(ctx, existingRelation.ID)
		if err != nil {
			return err
		}

		return nil
	}

	err = s.CreateCommentEmote(ctx, commentID, emote)
	if err != nil {
		return err
	}

	return nil
}

// CreateCommentEmote add an emote to an comment
func (s *CommentService) CreateCommentEmote(ctx context.Context, commentID string, emote CommentRelationshipType) error {
	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	_, err := s.DBClient.RelateTo(ctx, userEmail, commentID, string(emote))
	return err
}

// DeleteCommentEmote deletes an emote previously creatd by a user
func (s *CommentService) DeleteCommentEmote(ctx context.Context, relationshipID string) error {
	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	emoteRelationship, err := s.DBClient.GetRelationship(ctx, relationshipID)
	if err != nil {
		return err
	}

	if emoteRelationship.In == userEmail {
		err = s.DBClient.DeleteObject(userEmail, relationshipID)
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("emote does not belong to user %s...no work was done", userEmail)
		return err
	}

	return nil
}

// RemoveComment soft-delete a comment based on its ID
func (s *CommentService) RemoveComment(ctx context.Context, commentID string) error {
	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)
	commentToDelete, err := s.GetCommentByID(ctx, commentID)
	if err != nil {
		return err
	}

	if !strings.EqualFold(commentToDelete.ControlFields.CreatedBy, userEmail) {
		err = fmt.Errorf(
			"comments can only be deleted by creating user...no work was done.  %s - %s",
			commentToDelete.ControlFields.CreatedBy,
			userEmail)

		return err
	}

	return s.DBClient.SoftDeleteObject(userEmail, commentToDelete)
}

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
