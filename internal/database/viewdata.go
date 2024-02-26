package database

import (
	"bot-vk/internal/models"

	"gorm.io/gorm"
)

type ViewDataImpl struct {
	db *gorm.DB
}

func NewViewDataImpl(db *gorm.DB) *ViewDataImpl {
	return &ViewDataImpl{db: db}
}

func (impl *ViewDataImpl) AddView(dataId string, userId string, resource string) error {

	var View models.View

	View.UserId = userId
	View.DataId = dataId
	View.Resource = resource

	impl.db.Table("view").Create(&View)

	return nil
}
