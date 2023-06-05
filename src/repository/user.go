package repository

import (
	"github.com/peidrao/instago/src/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
}

type DBUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &DBUserRepository{db}
}

func (u *DBUserRepository) CreateUser(user *models.User) error {
	err := u.db.Create(&user).Error

	if err != nil {
		return err
	}
	return nil
}
