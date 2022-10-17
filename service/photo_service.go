package service

import (
	"github.com/hrswcksono/mygram-hacktiv/dto"
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/repository/photo_repository"
)

type PhotoService interface {
	AddPhoto(dto *dto.CreatePhotoRequest, userId int) (*dto.CreatedPhotoResponse, error)
	ReadAllPhoto() ([]dto.GetPhotoResponse, error)
	EditPhoto(edit *dto.EditPhotoRequest, photoId int) (*dto.EditPhotoResponse, error)
	DeletePhoto(photoId int) error
}

type photoService struct {
	photoRepo photo_repository.PhotoRepository
}

func NewPhotoService(photoRepo photo_repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepo: photoRepo,
	}
}

func (p *photoService) AddPhoto(payload *dto.CreatePhotoRequest, userId int) (*dto.CreatedPhotoResponse, error) {
	photoRequest := &entity.Photo{}

	payload.CreatePhotoRequestMapper(photoRequest)
	photoRequest.UserID = userId

	data, err := p.photoRepo.CreatePhoto(photoRequest)

	if err != nil {
		return nil, err
	}

	response := &dto.CreatedPhotoResponse{
		ID:        data.ID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoUrl:  data.PhotoUrl,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
	}

	return response, nil
}

func (p *photoService) ReadAllPhoto() ([]dto.GetPhotoResponse, error) {
	photo, err := p.photoRepo.GetAllPhoto()

	if err != nil {
		return nil, err
	}

	return dto.MapperResponseGet(photo), nil
}

func (p *photoService) EditPhoto(edit *dto.EditPhotoRequest, userId int) (*dto.EditPhotoResponse, error) {

	editPhotoRequest := &entity.Photo{}

	edit.EditPhotoRequestMapper(editPhotoRequest)

	editPhotoRequest.ID = userId

	photo, err := p.photoRepo.UpdatePhoto(editPhotoRequest)

	if err != nil {
		return nil, err
	}

	response := &dto.EditPhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	return response, nil
}

func (p *photoService) DeletePhoto(photoId int) error {
	err := p.photoRepo.DeletePhoto(photoId)
	if err != nil {
		return err
	}
	return nil
}
