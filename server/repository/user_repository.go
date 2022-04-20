package repository

import (
	"server/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindOneByUsername(username string) *model.User
}

type UserRepository struct {
	DB *gorm.DB
}

func CreateUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (repo *UserRepository) FindOneByUsername(username string) *model.User {
	return &model.User{}
}
