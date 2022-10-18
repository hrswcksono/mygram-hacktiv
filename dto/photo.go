package dto

import (
	"time"

	"github.com/hrswcksono/mygram-hacktiv/entity"
)

type CreatePhotoRequest struct {
	Title    string `json:"title" example:"pegunungan"`
	Caption  string `json:"caption" example:"gunung yang ada di indonesia"`
	PhotoUrl string `json:"photo_url" example:"google.com"`
}

type CreatedPhotoResponse struct {
	ID        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"pegunungan"`
	Caption   string    `json:"caption" example:"gunung yang ada di indonesia"`
	PhotoUrl  string    `json:"photo_url" example:"google.com"`
	UserID    int       `json:"user_id" example:"2"`
	CreatedAt time.Time `json:"created_at" example:"2022-10-17T13:44:00.294125+07:00"`
}

type GetPhotoResponse struct {
	ID        int          `json:"id" example:"1"`
	Title     string       `json:"title" example:"pegunungan"`
	Caption   string       `json:"caption" example:"gunung yang ada di indonesia"`
	PhotoUrl  string       `json:"photo_url" example:"google.com"`
	UserID    int          `json:"user_id" example:"2"`
	CreatedAt time.Time    `json:"created_at" example:"2022-10-17T13:44:00.294125+07:00"`
	UpdatedAt time.Time    `json:"updated_at" example:"2022-10-17T13:44:00.294125+07:00"`
	User      GetPhotoUser `json:"User"`
}

type GetPhotoUser struct {
	Email    string `json:"email" example:"email@email.com"`
	Username string `json:"username" example:"username1"`
}

type EditPhotoRequest struct {
	Title    string `json:"title" example:"pegunungan"`
	Caption  string `json:"caption" example:"gunung yang ada di indonesia"`
	PhotoUrl string `json:"photo_url" example:"google.com"`
}

type EditPhotoResponse struct {
	ID        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"pegunungan"`
	Caption   string    `json:"caption" example:"gunung yang ada di indonesia"`
	PhotoUrl  string    `json:"photo_url" example:"google.com"`
	UserID    int       `json:"user_id" example:"2"`
	UpdatedAt time.Time `json:"updated_at" example:"2022-10-17T13:44:00.294125+07:00"`
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
