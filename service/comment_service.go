package service

import (
	"github.com/hrswcksono/mygram-hacktiv/dto"
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/repository/comment_repository"
)

type CommentService interface {
	AddComment(dtoCreate *dto.CreateCommentRequest, userId int) (*dto.CreateCommentResponse, error)
	ReadAllComment() ([]dto.GetCommentResponse, error)
	EditComment(dtoEdit *dto.EditCommentRequest, userId int) (*dto.EditCommentResponse, error)
	DeleteComment(commentId int) error
}

type commentService struct {
	commentRepo comment_repository.CommentRepository
}

func NewCommentService(commentRepo comment_repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (c *commentService) AddComment(dtoCreate *dto.CreateCommentRequest, userId int) (*dto.CreateCommentResponse, error) {
	commentRequest := &entity.Comment{
		UserID:  userId,
		Message: dtoCreate.Message,
		PhotoID: dtoCreate.PhotoID,
	}

	data, err := c.commentRepo.CreateComment(commentRequest)

	if err != nil {
		return nil, err
	}

	response := &dto.CreateCommentResponse{
		ID:        data.ID,
		Message:   data.Message,
		PhotoID:   data.PhotoID,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
	}

	return response, nil
}

func (c *commentService) ReadAllComment() ([]dto.GetCommentResponse, error) {
	data, err := c.commentRepo.GetAllComment()

	if err != nil {
		return nil, err
	}

	return dto.MapperCommentGetResponse(data), nil
}

func (c *commentService) EditComment(dtoEdit *dto.EditCommentRequest, userId int) (*dto.EditCommentResponse, error) {
	editRequest := &entity.Comment{
		ID:      userId,
		Message: dtoEdit.Message,
	}

	print(editRequest)

	comment, err := c.commentRepo.UpdateComment(editRequest)

	if err != nil {
		return nil, err
	}

	response := &dto.EditCommentResponse{
		ID:        comment.ID,
		Title:     comment.Photo.Title,
		Caption:   comment.Photo.Caption,
		PhotoUrl:  comment.Photo.PhotoUrl,
		UserID:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}

	return response, nil
}

func (c *commentService) DeleteComment(commentId int) error {
	err := c.commentRepo.DeleteComment(commentId)
	if err != nil {
		return err
	}
	return nil
}
