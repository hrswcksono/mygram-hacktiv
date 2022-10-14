package entity

import (
	"time"
)

type User struct {
	ID           int
	Username     string
	Email        string
	Password     string
	Age          int
	Photos       Photo
	Comments     Comment
	SocialMedias SocialMedia
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	_, errCreate := govalidator.ValidateStruct(u)

// 	if errCreate != nil {
// 		err = errCreate
// 		return
// 	}

// 	u.Password = auth_helper.HashPass(u.Password)
// 	err = nil
// 	return
// }
