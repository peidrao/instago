package repository

import (
	"github.com/peidrao/instago/src/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	RemoveUser(id uint64) error
	LastUser() (*models.User, error)
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

func (u *DBUserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, id).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *DBUserRepository) RemoveUser(userID uint64) error {
	err := u.db.Delete(&models.User{}, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *DBUserRepository) LastUser() (*models.User, error) {
	var user models.User

	err := u.db.Last(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil

}
