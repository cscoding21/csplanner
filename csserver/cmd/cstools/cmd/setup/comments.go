package setup

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/services/comment"
	"csserver/internal/utils/quilljs"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func CreateTestComments(ctx context.Context) error {
	service := factory.GetCommentService()
	projectID := "project:1"

	results, err := service.FindProjectComments(ctx, projectID)
	if err != nil {
		log.Errorf("error checking existing comments %v", err)
		return err
	}

	if *results.Pagination.TotalResults > 0 {
		return nil
	}

	commentDelta := quilljs.New(nil)
	attr := make(map[string]interface{})
	attr["bold"] = true
	commentDelta.Insert("Here is my test comment", attr)
	js, err := json.Marshal(commentDelta)
	if err != nil {
		log.Errorf("error marshalling comment delta: %v", err)
		return err
	}

	comments := []comment.Comment{
		{
			ControlFields: common.ControlFields{
				ID: "comment:1",
			},
			ProjectID: projectID,
			Text:      string(js),
			IsEdited:  false,
		},
	}

	for _, c := range comments {
		_, err := service.AddComment(ctx, c)
		if err != nil {
			log.Errorf("error adding comment: %v", err)
			return err
		}

		//log.Infof("comment: %v", result)
	}

	replyDelta := quilljs.New(nil)
	replyDelta.Insert("Here is my test reply", attr)
	js, err = json.Marshal(replyDelta)
	if err != nil {
		log.Errorf("error marshalling reply delta: %v", err)
		return err
	}

	replies := []comment.Comment{
		{
			ProjectID: projectID,
			Text:      string(js),
			IsEdited:  false,
		},
	}

	for _, r := range replies {
		_, err := service.AddCommentReply(ctx, r, "comment:1")
		if err != nil {
			log.Errorf("error adding reply: %v", err)
			return err
		}

		//log.Infof("reply: %v", result)
	}

	return nil
}
