package service

import "bot-vk/internal/database"

type ViewServiceImpl struct {
	ViewData *database.DataBase
}

func NewViewServiceImpl(database *database.DataBase) *ViewServiceImpl {
	return &ViewServiceImpl{ViewData: database}
}

func (impl *ViewServiceImpl) AddView(dataId string, userId string, resource string) error {
	impl.ViewData.AddView(dataId, userId, resource)
	return nil
}
