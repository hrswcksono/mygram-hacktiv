package social_media_pg

import (
	"github.com/hrswcksono/mygram-hacktiv/entity"
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

func (s *smediaPG) CreateSocialMedia(smedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&smedia).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return smedia, tx.Commit().Error
}

func (s *smediaPG) GetAllSocialMedia() ([]entity.SocialMedia, error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	var smedia = []entity.SocialMedia{}

	if err := tx.Preload("User", func(db *gorm.DB) *gorm.DB { return db.Find(&entity.User{}) }).Find(&smedia).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return smedia, tx.Commit().Error
}

func (s *smediaPG) UpdateSocialMedia(smedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Updates(smedia).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return smedia, tx.Commit().Error
}

func (s *smediaPG) DeleteSocialMedia(smediaId int) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var smedia = &entity.SocialMedia{}

	if err := tx.Delete(&smedia, smediaId).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *smediaPG) GetSocialMediaByID(smediaId int) (*entity.SocialMedia, error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	var smedia = &entity.SocialMedia{}

	if err := tx.Find(&smedia, smediaId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return smedia, tx.Commit().Error
}
