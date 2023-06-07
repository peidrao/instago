package interfaces

import "github.com/peidrao/instago/src/domain/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByID(id uint64) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	FindLastUser() (*models.User, error)
	ListAllUsers() ([]*models.User, error)
	DestroyUser(id uint64) error
	UpdateUser(*models.User) (*models.User, error)
}
