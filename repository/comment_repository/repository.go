package comment_repository

import "github.com/hrswcksono/mygram-hacktiv/entity"

type CommentRepository interface {
	CreateComment(commentPayload *entity.Comment) (*entity.Comment, error)
	GetAllComment() ([]entity.Comment, error)
	UpdateComment(commentPayload *entity.Comment) (*entity.Comment, error)
	DeleteComment(commentId int) error
	GetCommentByID(commentId int) (*entity.Comment, error)
}
