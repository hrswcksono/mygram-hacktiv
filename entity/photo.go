package entity

import (
	"time"
)

type Photo struct {
	ID        int
	Title     string `gorm:"not null"`
	Caption   string
	PhotoUrl  string `gorm:"not null"`
	UserID    int
	Comments  Comment
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User
}

// func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
// 	_, errUpdate := govalidator.ValidateStruct(u)

// 	if errUpdate != nil {
// 		err = errUpdate
// 		return
// 	}
// 	return
// }

// func (u *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
// 	_, errUpdate := govalidator.ValidateStruct(u)

// 	if errUpdate != nil {
// 		err = errUpdate
// 		return
// 	}
// 	return
// }
