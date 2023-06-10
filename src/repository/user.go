package repository

import (
	"github.com/peidrao/instago/src/domain/interfaces"
	"github.com/peidrao/instago/src/domain/models"
	"gorm.io/gorm"
)

type DBUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserInterface {
	return &DBUserRepository{db}
}

func (u *DBUserRepository) CreateUser(user *models.User) error {
	err := u.db.Create(&user).Error

	if err != nil {
		return err
	}
	return nil
}

func (u *DBUserRepository) FindUserByID(id uint) (*models.User, error) {
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

func (u *DBUserRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := u.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *DBUserRepository) ListAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := u.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
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

func (u *DBUserRepository) FollowUser(userID, followerID uint) error {

	follower, err := u.FindUserByID(userID)

	if err != nil {
		return err
	}

	followed, err := u.FindUserByID(followerID)
	if err != nil {
		return err
	}

	follower.Following = append(follower.Following, followed)

	_, err = u.UpdateUser(follower)

	return err
}

func (u *DBUserRepository) FindFollowers(username string) ([]*models.User, error) {
	var user models.User
	result := u.db.Preload("Followers").Preload("Following").Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user.Followers, nil
}

func (u *DBUserRepository) FindFollowings(username string) ([]*models.User, error) {
	var user models.User
	result := u.db.Preload("Followers").Preload("Following").Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user.Following, nil
}
