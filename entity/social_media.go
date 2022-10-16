package entity

import "time"

type SocialMedia struct {
	ID             int
	Name           string `gorm:"not null"`
	SocialMediaURL string `gorm:"not null"`
	UserID         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           *User
}
