package service

import "bot-vk/internal/database"

type UserServiceImpl struct {
	UserData *database.DataBase
}

func NewUserServiceImpl(database *database.DataBase) *UserServiceImpl {
	return &UserServiceImpl{UserData: database}
}

func (impl *UserServiceImpl) AddUser(userId int64, firstName string, lastName string, userName string, resource string) error {
	impl.UserData.AddUser(userId, firstName, lastName, userName, resource)
	return nil
}
