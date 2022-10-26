package service

import (
	"fmt"

	"github.com/hrswcksono/mygram-hacktiv/dto"
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/auth_helper"
	"github.com/hrswcksono/mygram-hacktiv/pkg/errs"
	"github.com/hrswcksono/mygram-hacktiv/repository/user_repository"
)

type UserService interface {
	Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	Login(user *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	UpdateUser(user *dto.UserEditRequest, userId int) (*dto.UserEditResponse, errs.MessageErr)
	DeleteUser(userId int) errs.MessageErr
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {
	registerRequest := &entity.User{}

	fmt.Println(userPayload)
	userPayload.RegisterRequestMapper(registerRequest)

	data, err := u.userRepo.RegisterUser(registerRequest)

	if err != nil {
		return nil, err
	}

	return dto.ToRegisterResponse(*data), nil
}

func (u *userService) Login(user *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {

	userPayload := &entity.User{
		Email: user.Email,
	}

	userData, err := u.userRepo.LoginUser(userPayload)

	if err != nil {
		return nil, err
	}

	validPassword := auth_helper.ComparePass([]byte(userData.Password), []byte(user.Password))

	if !validPassword {
		return nil, errs.NewNotAuthenticated("error compare password")
	}

	token := auth_helper.GenerateToken(uint(userData.ID), userData.Username)

	response := &dto.LoginResponse{
		Token: token,
	}

	return response, nil
}

func (u *userService) UpdateUser(user *dto.UserEditRequest, userId int) (*dto.UserEditResponse, errs.MessageErr) {
	var editRequest = &entity.User{}

	user.EditRequestMapper(editRequest)

	data, err := u.userRepo.UpdateUser(editRequest, userId)

	if err != nil {
		return nil, err
	}

	response := &dto.UserEditResponse{
		ID:        data.ID,
		Email:     data.Email,
		Username:  data.Username,
		Age:       data.Age,
		UpdatedAt: data.UpdatedAt,
	}

	return response, nil
}

func (u *userService) DeleteUser(userId int) errs.MessageErr {
	err := u.userRepo.DeleteUser(userId)

	if err != nil {
		return err
	}
	return nil
}
