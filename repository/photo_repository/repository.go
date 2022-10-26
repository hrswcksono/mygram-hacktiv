package photo_repository

import (
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/errs"
)

type PhotoRepository interface {
	CreatePhoto(photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr)
	GetAllPhoto() ([]entity.Photo, errs.MessageErr)
	UpdatePhoto(photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr)
	DeletePhoto(photoId int) errs.MessageErr
	GetPhotoByID(photoId int) (*entity.Photo, errs.MessageErr)
}
