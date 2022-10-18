package dto

import (
	"time"

	"github.com/hrswcksono/mygram-hacktiv/entity"
)

type CreateSMediaRequest struct {
	Name           string `json:"name" example:"twitter"`
	SocialMediaURL string `json:"social_media_url" example:"twitter.com"`
}

type CreateSMediaResponse struct {
	ID             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"twitter"`
	SocialMediaURL string    `json:"social_media_url" example:"twitter.com"`
	UserID         int       `json:"user_id" example:"2"`
	CreatedAt      time.Time `json:"created_at" example:"2022-10-17T13:44:00.294125+07:00"`
}

type GetSMediaResponse struct {
	SocialMedia []SMediaList `json:"social_medias"`
}

type SMediaList struct {
	ID             int        `json:"id" example:"1"`
	Name           string     `json:"name" example:"twitter"`
	SocialMediaURL string     `json:"social_media_url" example:"twitter.com"`
	UserID         int        `json:"user_id" example:"2"`
	CreatedAt      time.Time  `json:"created_at" example:"2022-10-17T13:44:00.294125+07:00"`
	UpdatedAt      time.Time  `json:"updated_at" example:"2022-10-17T13:44:00.294125+07:00"`
	User           SMediaUser `json:"User"`
}

type SMediaUser struct {
	ID       int    `json:"id" example:"1"`
	Username string `json:"username" example:"username1"`
}

type EditSMediaRequest struct {
	Name           string `json:"name" example:"twitter"`
	SocialMediaURL string `json:"social_media_url" example:"twitter.com"`
}

type EditSMediaResponse struct {
	ID             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"twitter"`
	SocialMediaURL string    `json:"social_media_url" example:"twitter.com"`
	UserID         int       `json:"user_id" example:"2"`
	UpdatedAt      time.Time `json:"updated_at" example:"2022-10-17T13:44:00.294125+07:00"`
}

func toGetListSMediaResponse(in entity.SocialMedia) *SMediaList {
	return &SMediaList{
		ID:             in.ID,
		Name:           in.Name,
		SocialMediaURL: in.SocialMediaURL,
		UserID:         in.UserID,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
		User: SMediaUser{
			ID:       in.User.ID,
			Username: in.User.Username,
		},
	}
}

func MapperListSMedia(input []entity.SocialMedia) []SMediaList {
	var output []SMediaList
	for _, v := range input {
		newItem := *toGetListSMediaResponse(v)
		output = append(output, newItem)
	}
	return output
}
