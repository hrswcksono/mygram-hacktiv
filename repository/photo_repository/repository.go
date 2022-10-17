package photo_repository

import "github.com/hrswcksono/mygram-hacktiv/entity"

type PhotoRepository interface {
	CreatePhoto(photoPayload *entity.Photo) (*entity.Photo, error)
	GetAllPhoto() ([]entity.Photo, error)
	UpdatePhoto(photoPayload *entity.Photo) (*entity.Photo, error)
	DeletePhoto(photoId int) error
	GetPhotoByID(photoId int) (*entity.Photo, error)
}
