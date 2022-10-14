package user_pg

import (
	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/repository/user_repository"
	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) LoginUser(userPayload *entity.User) (*entity.User, error) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(userPayload).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var user = &entity.User{}

	if err := tx.Find(&user).Where("email = ?", userPayload.Email).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, tx.Commit().Error
}
