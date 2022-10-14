package dto_request

import (
	"github.com/asaskevich/govalidator"
	"github.com/hrswcksono/mygram-hacktiv/entity"
)

type RegisterRequest struct {
	Age      int
	Email    string
	Password string
	Username string
}

func (input RegisterRequest) RegisterRequestMapper(output *entity.User) {
	output.Age = input.Age
	output.Password = input.Password
	output.Username = input.Username
}

func (m *RegisterRequest) Validate() error {
	_, err := govalidator.ValidateStruct(m)

	if err != nil {
		return err
	}

	return nil
}
