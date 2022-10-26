package social_media_pg

import (
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/errs"
	"github.com/hrswcksono/mygram-hacktiv/repository/social_media_repository"
	"gorm.io/gorm"
)

type smediaPG struct {
	db *gorm.DB
}

func NewSMediaPG(db *gorm.DB) social_media_repository.SocialMediaRepository {
	return &smediaPG{
		db: db,
	}
}

func (s *smediaPG) CreateSocialMedia(smedia *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Create(&smedia).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return smedia, nil
}

func (s *smediaPG) GetAllSocialMedia() ([]entity.SocialMedia, errs.MessageErr) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	var smedia = []entity.SocialMedia{}

	if err := tx.Preload("User", func(db *gorm.DB) *gorm.DB { return db.Find(&entity.User{}) }).Find(&smedia).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return smedia, nil
}

func (s *smediaPG) UpdateSocialMedia(smedia *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Updates(smedia).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return smedia, nil
}

func (s *smediaPG) DeleteSocialMedia(smediaId int) errs.MessageErr {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return errs.NewInternalServerErrorr("something went wrong")
	}

	var smedia = &entity.SocialMedia{}

	if err := tx.Delete(&smedia, smediaId).Error; err != nil {
		tx.Rollback()
		return errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return errs.NewInternalServerErrorr("something went wrong")
	}

	return nil
}

func (s *smediaPG) GetSocialMediaByID(smediaId int) (*entity.SocialMedia, errs.MessageErr) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	var smedia = &entity.SocialMedia{}

	if err := tx.Find(&smedia, smediaId).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return smedia, nil
}
