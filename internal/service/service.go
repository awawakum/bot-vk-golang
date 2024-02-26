package service

import (
	"bot-vk/internal/database"
	"bot-vk/internal/models"
)

type UserService interface {
	AddUser(userId int64, firstName string, lastName string, userName string, resource string) error
}

type CategoryService interface {
	GetCategories() []models.Category
}

type KeyboardService interface {
	GetMainKeyboard() (models.Keyboard, error)
	GetCategoriesInlineKeyboard() (models.Keyboard, error)
	GetProductInlineKeyboard(category string) (models.LinkKeyboard, error)
	GetOSAGOInlineKeyboard() (models.LinkKeyboard, error)
	GetHistoryInlineKeyboard() (models.LinkKeyboard, error)
	GetBankInlineKeyboard() (models.LinkKeyboard, error)
}

type ViewService interface {
	AddView(dataId string, userId string, resource string) error
}

type Service struct {
	UserService
	CategoryService
	KeyboardService
	ViewService
}

func NewService(database *database.DataBase) *Service {
	return &Service{
		UserService:     NewUserServiceImpl(database),
		CategoryService: NewCategoryServiceImpl(database),
		KeyboardService: NewKeyboardServiceImpl(database),
		ViewService:     NewViewServiceImpl(database),
	}
}
