package entity

import (
	"time"
)

type Photo struct {
	ID        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    int
	Comments  Comment
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User
}
