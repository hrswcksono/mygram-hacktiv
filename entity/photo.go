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
