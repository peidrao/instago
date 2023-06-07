package repository

import (
	"github.com/peidrao/instago/src/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByID(id uint64) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindLastUser() (*models.User, error)
	DestroyUser(id uint64) error
	UpdateUser(*models.User) (*models.User, error)
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

func (u *DBUserRepository) FindUserByID(id uint64) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, id).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *DBUserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *DBUserRepository) DestroyUser(userID uint64) error {
	err := u.db.Delete(&models.User{}, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *DBUserRepository) FindLastUser() (*models.User, error) {
	var user models.User

	err := u.db.Last(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *DBUserRepository) UpdateUser(user *models.User) (*models.User, error) {
	err := u.db.Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
