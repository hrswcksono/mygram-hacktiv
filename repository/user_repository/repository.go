package user_repository

import "github.com/hrswcksono/mygram-hacktiv/entity"

type UserRepository interface {
	// RegiterUser()
	LoginUser(user *entity.User) (*entity.User, error)
	// UpdateUser() // param buat apa
	// DeleteUser()
}
