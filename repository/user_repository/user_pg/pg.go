package user_pg

import (
	"fmt"

	"github.com/hrswcksono/mygram-hacktiv/entity"
	"github.com/hrswcksono/mygram-hacktiv/pkg/errs"
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

func (u *userPG) RegisterUser(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Create(userPayload).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	var user = &entity.User{}

	if err := tx.Where("email = ?", userPayload.Email).Find(&user).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return user, nil
}

func (u *userPG) LoginUser(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	var user = &entity.User{}

	if err := tx.Where("email = ?", userPayload.Email).Find(&user).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return user, nil
}

func (u *userPG) UpdateUser(userPayload *entity.User, userId int) (*entity.User, errs.MessageErr) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	userPayload.ID = userId

	fmt.Println(userPayload)
	if err := tx.Updates(userPayload).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	var user = &entity.User{}

	if err := tx.First(&user, userId).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return user, nil
}

func (u *userPG) GetUserByID(userId int) (*entity.User, errs.MessageErr) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	var user = &entity.User{}

	if err := tx.First(&user, userId).Error; err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errs.NewInternalServerErrorr("something went wrong")
	}

	return user, nil
}

func (u *userPG) DeleteUser(userId int) errs.MessageErr {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return errs.NewInternalServerErrorr("something went wrong")
	}

	var user = &entity.User{}

	if err := tx.Delete(&user, userId).Error; err != nil {
		tx.Rollback()
		return errs.NewInternalServerErrorr("something went wrong")
	}

	if err := tx.Commit().Error; err != nil {
		return errs.NewInternalServerErrorr("something went wrong")
	}

	return nil
}
