package augment

import "csserver/internal/appserv/graph/idl"

func AugmentComment(comment *idl.Comment) {
	user := getUserByEmail(comment.CreatedBy)
	comment.User = user

	if comment.Replies != nil {
		AugmentCommentSlice(&comment.Replies)
	}
}

func AugmentCommentSlice(comments *[]*idl.Comment) {
	for _, comment := range *comments {
		AugmentComment(comment)
	}
}
