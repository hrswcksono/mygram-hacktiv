package user_repository

import (
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/errs"
)

type UserRepository interface {
	LoginUser(user *entity.User) (*entity.User, errs.MessageErr)
	RegisterUser(user *entity.User) (*entity.User, errs.MessageErr)
	UpdateUser(user *entity.User, userId int) (*entity.User, errs.MessageErr)
	GetUserByID(userId int) (*entity.User, errs.MessageErr)
	DeleteUser(userId int) errs.MessageErr
}
