package models

import (
	"time"
)

type User struct {
	UserId    int64
	FirstName string
	LastName  string
	UserName  string
	Resource  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
