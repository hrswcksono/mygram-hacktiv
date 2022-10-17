package social_media_repository

import "github.com/hrswcksono/mygram-hacktiv/entity"

type SocialMediaRepository interface {
	CreateSocialMedia(smedia *entity.SocialMedia) (*entity.SocialMedia, error)
	GetAllSocialMedia() ([]entity.SocialMedia, error)
	UpdateSocialMedia(smedia *entity.SocialMedia) (*entity.SocialMedia, error)
	DeleteSocialMedia(smediaId int) error
	GetSocialMediaByID(smediaId int) (*entity.SocialMedia, error)
}
