package database

import (
	"bot-vk/internal/models"

	"gorm.io/gorm"
)

type UserData interface {
	AddUser(userId int64, firstName string, lastName string, userName string, resource string) error
}

type CategoryData interface {
	GetCategories() []models.Category
}

type PorductData interface {
	GetProductsWhereCategory(category string) []models.Product
}

type ViewData interface {
	AddView(dataId string, userId string, resource string) error
}

type DataBase struct {
	UserData
	CategoryData
	PorductData
	ViewData
}

func NewDataBase(db *gorm.DB) *DataBase {
	return &DataBase{
		UserData:     NewUserDataImpl(db),
		CategoryData: NewCategoryDataImpl(db),
		PorductData:  NewProductDataImpl(db),
		ViewData:     NewViewDataImpl(db),
	}
}
