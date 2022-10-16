package user_repository

import "github.com/hrswcksono/mygram-hacktiv/entity"

type UserRepository interface {
	LoginUser(user *entity.User) (*entity.User, error)
	RegisterUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User, userId int) (*entity.User, error)
	GetUserByID(userId int) (*entity.User, error)
	DeleteUser(userId int) error
}
