package database

import (
	"bot-vk/internal/models"

	"gorm.io/gorm"
)

type UserDataImpl struct {
	db *gorm.DB
}

func NewUserDataImpl(db *gorm.DB) *UserDataImpl {
	return &UserDataImpl{db: db}
}

func (impl *UserDataImpl) AddUser(userId int64, firstName string, lastName string, userName string, resource string) error {

	var User models.User

	User.UserId = userId

	User.FirstName = firstName
	User.LastName = lastName
	User.UserName = userName
	User.Resource = resource

	impl.db.Table("user").Select("user_id", "first_name", "last_name", "username", "resource").Create(&User)

	return nil
}
