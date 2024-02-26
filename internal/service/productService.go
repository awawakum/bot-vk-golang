package service

import (
	"bot-vk/internal/database"
)

type ProductServiceImpl struct {
	ProductData *database.DataBase
}

func NewProductServiceImpl(database *database.DataBase) *ProductServiceImpl {
	return &ProductServiceImpl{ProductData: database}
}
