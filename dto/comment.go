package dto

import (
	"time"

	"github.com/hrswcksono/mygram-hacktiv/entity"
)

type CreateCommentRequest struct {
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
}

type CreateCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetCommentResponse struct {
	ID        int          `json:"id"`
	Message   string       `json:"message"`
	PhotoID   int          `json:"photo_id"`
	UserID    int          `json:"user_id"`
	UpdatedAt time.Time    `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
	User      UserComment  `json:"User"`
	Photo     PhotoComment `json:"Photo"`
}

type UserComment struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoComment struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

type EditCommentRequest struct {
	Message string `json:"message"`
}

type EditCommentResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToGetCommentResponse(data entity.Comment) *GetCommentResponse {
	return &GetCommentResponse{
		ID:        data.ID,
		Message:   data.Message,
		PhotoID:   data.PhotoID,
		UserID:    data.UserID,
		UpdatedAt: data.UpdatedAt,
		CreatedAt: data.CreatedAt,
		User: UserComment{
			ID:       data.User.ID,
			Email:    data.User.Email,
			Username: data.User.Username,
		},
		Photo: PhotoComment{
			ID:       data.Photo.ID,
			Title:    data.Photo.Title,
			PhotoUrl: data.Photo.PhotoUrl,
			UserID:   data.Photo.UserID,
		},
	}
}

func MapperCommentGetResponse(input []entity.Comment) []GetCommentResponse {
	var output []GetCommentResponse
	for _, v := range input {
		newItem := *ToGetCommentResponse(v)
		output = append(output, newItem)
	}
	return output
}
