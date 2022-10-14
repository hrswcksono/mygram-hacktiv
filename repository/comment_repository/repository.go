package comment_repository

import "github.com/hrswcksono/mygram-hacktiv/entity"

type CommentRepository interface {
	CreateComment(commentPayload entity.Comment)
	GetAllComment()
	UpdateComment(commentId int)
	DeleteComment(commentId int)
}
