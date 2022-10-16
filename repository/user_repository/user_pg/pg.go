package user_pg

import (
	"errors"
	"fmt"

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

func (u *userPG) RegisterUser(userPayload *entity.User) (*entity.User, error) {
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

	if err := tx.Where("email = ?", userPayload.Email).Find(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, tx.Commit().Error
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

	var user = &entity.User{}

	if err := tx.Where("email = ?", userPayload.Email).Find(&user).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("cannot get email/password")
	}

	return user, tx.Commit().Error
}

func (u *userPG) UpdateUser(userPayload *entity.User, userId int) (*entity.User, error) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	userPayload.ID = userId

	fmt.Println(userPayload)
	if err := tx.Updates(userPayload).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var user = &entity.User{}

	if err := tx.First(&user, userId).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed get user data")
	}

	return user, tx.Commit().Error
}

func (u *userPG) GetUserByID(userId int) (*entity.User, error) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	var user = &entity.User{}

	if err := tx.First(&user, userId).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed get user data")
	}

	return user, tx.Commit().Error
}

func (u *userPG) DeleteUser(userId int) error {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var user = &entity.User{}

	if err := tx.Delete(&user, userId).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
