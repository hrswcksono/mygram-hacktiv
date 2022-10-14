package entity

import (
	"time"
)

type Comment struct {
	ID        int
	UserID    int
	PhotoID   int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User
}
