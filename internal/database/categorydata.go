package database

import (
	"bot-vk/internal/models"

	"gorm.io/gorm"
)

type CategoryDataImpl struct {
	db *gorm.DB
}

func NewCategoryDataImpl(db *gorm.DB) *CategoryDataImpl {
	return &CategoryDataImpl{db: db}
}

func (impl *CategoryDataImpl) GetCategories() []models.Category {
	var categories []models.Category

	impl.db.Table("category").Find(&categories)

	return categories
}
