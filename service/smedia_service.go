package service

import (
	"github.com/hrswcksono/mygram-hacktiv/dto"
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/errs"
	"github.com/hrswcksono/mygram-hacktiv/repository/social_media_repository"
)

type SMediaService interface {
	AddSMedia(smedia *dto.CreateSMediaRequest, userId int) (*dto.CreateSMediaResponse, errs.MessageErr)
	ReadAllSMedia() (*dto.GetSMediaResponse, errs.MessageErr)
	EditSMedia(smedia *dto.EditSMediaRequest, smediaI int) (*dto.EditSMediaResponse, errs.MessageErr)
	DeleteSMedia(smediaId int) errs.MessageErr
}

type smediaService struct {
	smediaRepo social_media_repository.SocialMediaRepository
}

func NewSMediaService(smediaRepo social_media_repository.SocialMediaRepository) SMediaService {
	return &smediaService{
		smediaRepo: smediaRepo,
	}
}

func (s *smediaService) AddSMedia(smedia *dto.CreateSMediaRequest, userId int) (*dto.CreateSMediaResponse, errs.MessageErr) {
	smediaRequest := &entity.SocialMedia{
		Name:           smedia.Name,
		SocialMediaURL: smedia.SocialMediaURL,
		UserID:         userId,
	}

	data, err := s.smediaRepo.CreateSocialMedia(smediaRequest)

	if err != nil {
		return nil, err
	}

	response := &dto.CreateSMediaResponse{
		ID:             data.ID,
		Name:           data.Name,
		SocialMediaURL: data.SocialMediaURL,
		UserID:         data.UserID,
		CreatedAt:      data.CreatedAt,
	}

	return response, nil
}

func (s *smediaService) ReadAllSMedia() (*dto.GetSMediaResponse, errs.MessageErr) {
	smedia, err := s.smediaRepo.GetAllSocialMedia()

	if err != nil {
		return nil, err
	}

	response := &dto.GetSMediaResponse{
		SocialMedia: dto.MapperListSMedia(smedia),
	}

	return response, nil
}

func (s *smediaService) EditSMedia(smedia *dto.EditSMediaRequest, smediaI int) (*dto.EditSMediaResponse, errs.MessageErr) {

	editSMedia := &entity.SocialMedia{
		ID:             smediaI,
		Name:           smedia.Name,
		SocialMediaURL: smedia.SocialMediaURL,
	}

	data, err := s.smediaRepo.UpdateSocialMedia(editSMedia)

	if err != nil {
		return nil, err
	}

	response := &dto.EditSMediaResponse{
		ID:             data.ID,
		Name:           data.Name,
		SocialMediaURL: data.SocialMediaURL,
		UserID:         data.UserID,
		UpdatedAt:      data.UpdatedAt,
	}

	return response, nil
}

func (s *smediaService) DeleteSMedia(smediaId int) errs.MessageErr {
	err := s.smediaRepo.DeleteSocialMedia(smediaId)
	if err != nil {
		return err
	}
	return nil
}
