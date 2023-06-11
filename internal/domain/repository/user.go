package repository

import (
	"github.com/peidrao/instago/internal/domain/entity"
	"gorm.io/gorm"
)

type DBUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserInterface {
	return &DBUserRepository{db}
}

func (u *DBUserRepository) CreateUser(user *entity.User) error {
	err := u.db.Create(&user).Error

	if err != nil {
		return err
	}
	return nil
}

func (u *DBUserRepository) FindUserByID(id uint) (*entity.User, error) {
	var user entity.User
	err := u.db.First(&user, id).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *DBUserRepository) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *DBUserRepository) FindUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *DBUserRepository) ListAllUsers() ([]*entity.User, error) {
	var users []*entity.User
	err := u.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *DBUserRepository) DestroyUser(userID uint64) error {
	err := u.db.Delete(&entity.User{}, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *DBUserRepository) FindLastUser() (*entity.User, error) {
	var user entity.User

	err := u.db.Last(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *DBUserRepository) UpdateUser(user *entity.User) (*entity.User, error) {
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

func (u *DBUserRepository) FindFollowers(username string) ([]*entity.User, error) {
	var user entity.User
	result := u.db.Preload("Followers").Preload("Following").Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user.Followers, nil
}

func (u *DBUserRepository) FindFollowing(username string) ([]*entity.User, error) {
	var user entity.User
	result := u.db.Preload("Followers").Preload("Following").Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user.Following, nil
}
