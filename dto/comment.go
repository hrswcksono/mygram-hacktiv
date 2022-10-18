package dto

import (
	"time"

	"github.com/hrswcksono/mygram-hacktiv/entity"
)

type CreateCommentRequest struct {
	Message string `json:"message" example:"message 1"`
	PhotoID int    `json:"photo_id" example:"1"`
}

type CreateCommentResponse struct {
	ID        int       `json:"id" example:"1"`
	Message   string    `json:"message" example:"message 1"`
	PhotoID   int       `json:"photo_id" example:"1"`
	UserID    int       `json:"user_id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2022-10-17T13:44:00.294125+07:00"`
}

type GetCommentResponse struct {
	ID        int          `json:"id" example:"1"`
	Message   string       `json:"message" example:"message 1"`
	PhotoID   int          `json:"photo_id" example:"1"`
	UserID    int          `json:"user_id" example:"1"`
	UpdatedAt time.Time    `json:"updated_at" example:"2022-10-17T13:44:00.294125+07:00"`
	CreatedAt time.Time    `json:"created_at" example:"2022-10-17T13:44:00.294125+07:00"`
	User      UserComment  `json:"User"`
	Photo     PhotoComment `json:"Photo"`
}

type UserComment struct {
	ID       int    `json:"id" example:"1"`
	Email    string `json:"email" example:"email@email.com"`
	Username string `json:"username" example:"username1"`
}

type PhotoComment struct {
	ID       int    `json:"id" example:"1"`
	Title    string `json:"title" example:"pegunungan"`
	Caption  string `json:"caption" example:"gunung yang ada di indonesia"`
	PhotoUrl string `json:"photo_url" example:"google.com"`
	UserID   int    `json:"user_id" example:"1"`
}

type EditCommentRequest struct {
	Message string `json:"message" example:"message 1"`
}

type EditCommentResponse struct {
	ID        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"pegunungan"`
	Caption   string    `json:"caption" example:"gunung yang ada di indonesia"`
	PhotoUrl  string    `json:"photo_url" example:"google.com"`
	UserID    int       `json:"user_id" example:"1"`
	UpdatedAt time.Time `json:"updated_at" example:"2022-10-17T13:44:00.294125+07:00"`
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
