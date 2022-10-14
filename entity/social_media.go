package entity

import "time"

type SocialMedia struct {
	ID             int
	Name           string
	SocialMediaURL string
	UserID         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           *User
}
