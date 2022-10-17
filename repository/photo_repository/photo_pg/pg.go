package photo_pg

import (
	"fmt"

	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/repository/photo_repository"
	"gorm.io/gorm"
)

type photoPG struct {
	db *gorm.DB
}

func NewPhotoPG(db *gorm.DB) photo_repository.PhotoRepository {
	return &photoPG{
		db: db,
	}
}

func (p *photoPG) CreatePhoto(photoPayload *entity.Photo) (*entity.Photo, error) {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&photoPayload).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return photoPayload, tx.Commit().Error
}

func (p *photoPG) GetAllPhoto() ([]entity.Photo, error) {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	var photo = []entity.Photo{}

	if err := tx.Preload("User", func(db *gorm.DB) *gorm.DB { return db.Find(&entity.User{}) }).Find(&photo).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return photo, tx.Commit().Error
}

func (p *photoPG) UpdatePhoto(photoPayload *entity.Photo) (*entity.Photo, error) {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	fmt.Println(photoPayload)

	if err := tx.Updates(photoPayload).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return photoPayload, tx.Commit().Error
}

func (p *photoPG) DeletePhoto(photoId int) error {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	fmt.Println(photoId)

	var photo = &entity.Photo{}

	if err := tx.Delete(&photo, photoId).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (p *photoPG) GetPhotoByID(photoId int) (*entity.Photo, error) {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	var photo = &entity.Photo{}

	if err := tx.Find(&photo, photoId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return photo, nil
}
