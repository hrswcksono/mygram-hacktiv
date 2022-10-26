package social_media_repository

import (
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/errs"
)

type SocialMediaRepository interface {
	CreateSocialMedia(smedia *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr)
	GetAllSocialMedia() ([]entity.SocialMedia, errs.MessageErr)
	UpdateSocialMedia(smedia *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr)
	DeleteSocialMedia(smediaId int) errs.MessageErr
	GetSocialMediaByID(smediaId int) (*entity.SocialMedia, errs.MessageErr)
}
