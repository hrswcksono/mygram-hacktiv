package service

import (
	"errors"
	"fmt"

	"github.com/hrswcksono/mygram-hacktiv/dto"
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/auth_helper"
	"github.com/hrswcksono/mygram-hacktiv/repository/user_repository"
)

type UserService interface {
	Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(user *dto.LoginRequest) (*dto.LoginResponse, error)
	UpdateUser(user *dto.UserEditRequest, userId int) (*dto.UserEditResponse, error)
	DeleteUser(userId int) error
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Register(userPayload *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	regiterRequest := &entity.User{}

	fmt.Println(userPayload)
	userPayload.RegisterRequestMapper(regiterRequest)

	data, err := u.userRepo.RegisterUser(regiterRequest)

	if err != nil {
		return nil, err
	}

	return dto.ToRegisterResponse(*data), nil
}

func (u *userService) Login(user *dto.LoginRequest) (*dto.LoginResponse, error) {

	userPayload := &entity.User{
		Email: user.Email,
	}

	userData, err := u.userRepo.LoginUser(userPayload)

	if err != nil {
		fmt.Println("error get user by email")
		return nil, err
	}

	validPassword := auth_helper.ComparePass([]byte(userData.Password), []byte(user.Password))

	if !validPassword {
		fmt.Println("error compare password")
		return nil, errors.New("invalid email/password")
	}

	token := auth_helper.GenerateToken(uint(userData.ID), userData.Username)

	response := &dto.LoginResponse{
		Token: token,
	}

	return response, nil
}

func (u *userService) UpdateUser(user *dto.UserEditRequest, userId int) (*dto.UserEditResponse, error) {
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

func (u *userService) DeleteUser(userId int) error {
	err := u.userRepo.DeleteUser(userId)

	if err != nil {
		return err
	}
	return nil
}
