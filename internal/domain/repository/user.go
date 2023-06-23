package repository

import (
	"log"

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

	log.Println(user)
	log.Println("FIND USER BY USERNAME", user.Username)

	if err := u.FindByAttr(user, entity.User{Username: username}); err != nil {
		return nil, err
	}

	log.Println("FIND USER BY USERNAME FINAL ->", user.Username)

	return user, nil
}

func (u *UserRepository) FindUserByID(ID uint) (*entity.User, error) {
	user := &entity.User{}
	attr := map[string]interface{}{"id": ID}

	if err := u.FindByAttr(user, attr); err != nil {
		return nil, err
	}
	return user, nil

}

func (u *UserRepository) FindUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	attr := map[string]interface{}{"email": email}

	if err := u.FindByAttr(user, attr); err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *UserRepository) UpdateUser(user *entity.User, ID uint) (*entity.User, error) {
	err := u.Update(entity.User{}, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
