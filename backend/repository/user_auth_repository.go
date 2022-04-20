package repository

import (
	"server/model"

	"gorm.io/gorm"
)

type IUserAuthRepository interface {
	FindOneById(ID uint) *model.User
	FindOneByUsername(Username string) *model.User
	Save(user model.User) (uint, error)
}

type UserAuthRepository struct {
	DB *gorm.DB
}

func CreateUserAuthRepository(DB *gorm.DB) *UserAuthRepository {
	return &UserAuthRepository{DB}
}

func (repo *UserAuthRepository) FindOneById(ID uint) *model.User {
	var userAuthData model.User
	repo.DB.Select("id", "username", "password").Find(&userAuthData, ID)
	return &userAuthData
}

func (repo *UserAuthRepository) FindOneByUsername(Username string) *model.User {
	var userAuthData model.User
	repo.DB.Select("id", "username", "password").Find(&userAuthData, "username = ?", Username)
	return &userAuthData
}

func (repo *UserAuthRepository) Save(user model.User) (uint, error) {
	result := repo.DB.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}
