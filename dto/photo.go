package dto

import (
	"time"

	"github.com/hrswcksono/mygram-hacktiv/entity"
)

type CreatePhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type CreatedPhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetPhotoResponse struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Caption   string       `json:"caption"`
	PhotoUrl  string       `json:"photo_url"`
	UserID    int          `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      GetPhotoUser `json:"User"`
}

type GetPhotoUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type EditPhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type EditPhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (input EditPhotoRequest) EditPhotoRequestMapper(output *entity.Photo) {
	output.Title = input.Title
	output.Caption = input.Caption
	output.PhotoUrl = input.PhotoUrl
}

func (input CreatePhotoRequest) CreatePhotoRequestMapper(output *entity.Photo) {
	output.Title = input.Title
	output.Caption = input.Caption
	output.PhotoUrl = input.PhotoUrl
}

func toGetPhotoResponse(in entity.Photo) *GetPhotoResponse {
	return &GetPhotoResponse{
		ID:        in.ID,
		Title:     in.Title,
		Caption:   in.Caption,
		PhotoUrl:  in.PhotoUrl,
		UserID:    in.UserID,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		User: GetPhotoUser{
			Email:    in.User.Email,
			Username: in.User.Username,
		},
	}
}

func MapperResponseGet(input []entity.Photo) []GetPhotoResponse {
	var output []GetPhotoResponse
	for _, v := range input {
		// var newItem GetPhotoResponse
		newItem := *toGetPhotoResponse(v)
		output = append(output, newItem)
	}
	return output
}
