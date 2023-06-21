package repository

import (
	"github.com/peidrao/instago/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	*GenericRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		GenericRepository: NewRepository(db),
	}
}

func (u *UserRepository) CreateUser(user *entity.User) error {
	return u.Create(user)
}

func (u *UserRepository) FindAllUsers() ([]entity.User, error) {
	var users []entity.User
	if err := u.FindAll(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) FindUserByUsername(username string) (*entity.User, error) {
	user := &entity.User{}
	if err := u.FindByAttr("username", username, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) FindUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	if err := u.FindByAttr("email", email, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) UpdateUser(user *entity.User, ID uint) (*entity.User, error) {
	err := u.Update(user, ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
