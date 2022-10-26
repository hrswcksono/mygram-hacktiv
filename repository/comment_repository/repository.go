package comment_repository

import (
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/errs"
)

type CommentRepository interface {
	CreateComment(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr)
	GetAllComment() ([]entity.Comment, errs.MessageErr)
	UpdateComment(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr)
	DeleteComment(commentId int) errs.MessageErr
	GetCommentByID(commentId int) (*entity.Comment, errs.MessageErr)
}
