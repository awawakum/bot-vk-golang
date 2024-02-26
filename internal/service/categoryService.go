package service

import (
	"bot-vk/internal/database"
	"bot-vk/internal/models"
)

type CategoryServiceImpl struct {
	CategoryData *database.DataBase
}

func NewCategoryServiceImpl(database *database.DataBase) *CategoryServiceImpl {
	return &CategoryServiceImpl{CategoryData: database}
}

func (impl *CategoryServiceImpl) GetCategories() []models.Category {
	categories := impl.CategoryData.GetCategories()
	return categories
}
