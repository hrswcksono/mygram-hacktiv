package dto

import (
	"time"

	"github.com/hrswcksono/mygram-hacktiv/entity"
)

type RegisterRequest struct {
	Age      int    `json:"age" valid:"required~age cannot be empty"`
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty"`
	Username string `json:"username" valid:"required~username cannot be empty"`
}

type RegisterResponse struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserEditRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserEditResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (input RegisterRequest) RegisterRequestMapper(output *entity.User) {
	output.Age = input.Age
	output.Password = input.Password
	output.Username = input.Username
	output.Email = input.Email
}

func (input LoginRequest) LoginRequestMapper(output *entity.User) {
	output.Email = input.Email
	output.Password = input.Password
}

func (input UserEditRequest) EditRequestMapper(output *entity.User) {
	output.Email = input.Email
	output.Username = input.Username
}

func ToRegisterResponse(in entity.User) *RegisterResponse {
	return &RegisterResponse{
		Age:      in.Age,
		Email:    in.Email,
		ID:       in.ID,
		Username: in.Username,
	}
}
