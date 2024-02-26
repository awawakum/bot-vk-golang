package models

import "time"

type Category struct {
	CategoryId    string `gorm:"primaryKey"`
	CategoryTitle string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
