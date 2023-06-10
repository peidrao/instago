package interfaces

import "github.com/peidrao/instago/src/domain/models"

type UserInterface interface {
	CreateUser(user *models.User) error
	FindUserByID(id uint) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	FindLastUser() (*models.User, error)
	ListAllUsers() ([]*models.User, error)
	DestroyUser(id uint64) error
	UpdateUser(*models.User) (*models.User, error)

	FollowUser(userId, followerID uint) error
	FindFollowers(username string) ([]*models.User, error)
	FindFollowing(username string) ([]*models.User, error)
}
