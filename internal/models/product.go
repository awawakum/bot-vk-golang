package models

import "time"

type Product struct {
	ProductId    string `gorm:"primaryKey"`
	ProductTitle string
	ProductBody  string
	ProductPhoto byte
	ProductTag   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
