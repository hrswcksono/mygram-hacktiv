package photo_pg

import (
	"fmt"

	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/errs"
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

func (p *photoPG) CreatePhoto(photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr) {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Create(&photoPayload).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return photoPayload, nil
}

func (p *photoPG) GetAllPhoto() ([]entity.Photo, errs.MessageErr) {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	var photo = []entity.Photo{}

	if err := tx.Preload("User", func(db *gorm.DB) *gorm.DB { return db.Find(&entity.User{}) }).Find(&photo).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return photo, nil
}

func (p *photoPG) UpdatePhoto(photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr) {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	fmt.Println(photoPayload)

	if err := tx.Updates(photoPayload).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return photoPayload, nil
}

func (p *photoPG) DeletePhoto(photoId int) errs.MessageErr {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return errs.NewInternalServerErrorr("something went wrong")
	}

	fmt.Println(photoId)

	var photo = &entity.Photo{}

	if err := tx.Delete(&photo, photoId).Error; err != nil {
		tx.Rollback()
		return errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return errs.NewInternalServerErrorr("something went wrong")
	}

	return nil
}

func (p *photoPG) GetPhotoByID(photoId int) (*entity.Photo, errs.MessageErr) {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	var photo = &entity.Photo{}

	if err := tx.Find(&photo, photoId).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return photo, nil
}
