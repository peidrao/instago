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

func (u *UserRepository) FindUserFollowersCount(user *entity.User) (uint, uint) {
	followers := u.DB.Model(user).Where("is_accept= true").Association("Following").Count()

	following := u.DB.Model(user).Where("is_accept= true").Association("Followers").Count()

	return uint(followers), uint(following)
}

func (u *UserRepository) FindUserByUsername(username string) (*entity.User, uint, uint, error) {
	user := &entity.User{}

	if err := u.FindByAttr(&user, entity.User{Username: username}); err != nil {
		return nil, 0, 0, err
	}

	followers, following := u.FindUserFollowersCount(user)

	return user, followers, following, nil
}

func (u *UserRepository) FindUserByID(ID uint) (*entity.User, uint, uint, error) {
	user := &entity.User{}
	attr := map[string]interface{}{"id": ID}

	if err := u.FindByAttr(&user, attr); err != nil {
		return nil, 0, 0, err
	}

	followers, following := u.FindUserFollowersCount(user)

	return user, followers, following, nil
}

func (u *UserRepository) FindUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	attr := map[string]interface{}{"email": email}

	if err := u.FindByAttr(&user, attr); err != nil {
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

func (u *UserRepository) FindUsersSuggestions(ID uint) []entity.User {
	var users []entity.User
	subQuery := u.DB.Table("follows").
		Select("following_id").
		Where("follower_id = ?", ID)

	u.DB.Table("users u").
		Where("u.id <> ?", ID).
		Not("u.id IN (?)", subQuery).
		Order("u.created_at ASC").
		Limit(3).
		Find(&users)

	return users
}
