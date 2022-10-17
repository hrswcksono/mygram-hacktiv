package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/hrswcksono/mygram-hacktiv/pkg/auth_helper"
	"gorm.io/gorm"
)

type User struct {
	ID           int
	Username     string      `gorm:"index:,unique, not null"`
	Email        string      `gorm:"index:,unique, not null" valid:"email"`
	Password     string      `gorm:"not null" valid:"minstringlength(6)"`
	Age          int         `gorm:"not null"`
	Photos       Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments     Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SocialMedias SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	if u.Age < 9 {
		return errors.New("age to young, min 9")
	}

	u.Password = auth_helper.HashPass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}
	return
}
