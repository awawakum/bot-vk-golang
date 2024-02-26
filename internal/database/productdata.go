package database

import (
	"bot-vk/internal/models"

	"gorm.io/gorm"
)

type ProductDataImpl struct {
	db *gorm.DB
}

func NewProductDataImpl(db *gorm.DB) *ProductDataImpl {
	return &ProductDataImpl{db: db}
}

func (impl *ProductDataImpl) GetProductsWhereCategory(category string) []models.Product {
	var products []models.Product

	impl.db.Table("product").Where("product_tag=?", category).Find(&products)

	return products
}
