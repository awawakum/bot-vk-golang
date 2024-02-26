package models

import (
	"time"
)

type View struct {
	DataId    string
	UserId    string
	Resource  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
