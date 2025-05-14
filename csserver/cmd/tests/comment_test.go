package tests

import (
	"csserver/internal/appserv/csmap"
	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/services/comment"
	"fmt"
	"testing"
)

func TestFindProjectComments(t *testing.T) {
	ctx := getTestContext()
	service := factory.GetCommentService(ctx)
	idUnderTest := "Tb6EXS4fTQ6kyEcZhKo_Sw"

	comments, err := service.FindProjectComments(ctx, idUnderTest)
	if err != nil {
		t.Error(err)
		return
	}

	for _, c := range comments.Results {
		base := csmap.GetDataEnvelope(&c)
		csmap.AugmentBaseModel(ctx, base)
		cut := csmap.CommentCommentToIdl(c.Data)
		fmt.Println(cut)

		drawCommentThread(c)
	}
}

func TestGetCommentThread(t *testing.T) {
	ctx := getTestContext()
	service := factory.GetCommentService(ctx)

	idUnderTest := "comment:1"

	comment, err := service.GetCommentThread(ctx, idUnderTest)
	if err != nil {
		t.Error(err)
		return
	}

	if comment != nil {
		drawCommentThread(*comment)
	} else {
		t.Error(fmt.Errorf("comment with id %s return null value", idUnderTest))
	}
}

func TestToggleCommentReaction(t *testing.T) {
	ctx := getTestContext()
	service := factory.GetCommentService(ctx)

	idUnderTest := "comment:1"
	projectUnderTest := "project:1"

	err := service.ToggleCommentReaction(ctx, projectUnderTest, idUnderTest, comment.Likes)
	if err != nil {
		t.Error(err)
	}
}

func drawCommentThread(item common.BaseModel[comment.Comment]) {
	fmt.Printf("%s : likes: %v | loves: %v | laughs: %v\n", item.Data.Text, item.Data.Likes, item.Data.Loves, item.Data.LaughsAt)

	if len(item.Data.CommentReplies) > 0 {
		for _, r := range item.Data.CommentReplies {
			fmt.Printf("  --- %v | loves: %v | laughs: %v : likes %v\n", r.Data.Text, r.Data.Likes, item.Data.Loves, item.Data.LaughsAt)
		}
	}
}
