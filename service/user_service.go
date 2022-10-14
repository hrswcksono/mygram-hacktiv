package service

import (
	"github.com/hrswcksono/mygram-hacktiv/dto/dto_request"
	"github.com/hrswcksono/mygram-hacktiv/repository/user_repository"
)

type UserService interface {
	UserLogin(userPayload *dto_request.RegisterRequest)
}

type userService struct {
	userRepo user_repository.UserRepository
}
