package dto

import (
	"time"

	"github.com/hrswcksono/mygram-hacktiv/entity"
)

type CreateSMediaRequest struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

type CreateSMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type GetSMediaResponse struct {
	SocialMedia []SMediaList `json:"social_medias"`
}

type SMediaList struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	SocialMediaURL string     `json:"social_media_url"`
	UserID         int        `json:"UserId"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	User           SMediaUser `json:"User"`
}

type SMediaUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type EditSMediaRequest struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

type EditSMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
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
