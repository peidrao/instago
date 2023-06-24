package repository

import (
	"github.com/peidrao/instago/internal/domain/entity"
	"gorm.io/gorm"
)

type FollowRepository struct {
	*GenericRepository
}

func NewFollowRepository(db *gorm.DB) *FollowRepository {
	return &FollowRepository{
		GenericRepository: NewRepository(db),
	}
}

func (f *FollowRepository) CreateFollow(follow *entity.Follow) error {
	return f.Create(follow)
}

func (f *FollowRepository) UpdateFollow(follow *entity.Follow, attr interface{}) error {
	return f.Update(follow, attr)
}

func (f *FollowRepository) FindLinkFollows(attr interface{}) bool {
	follow := &entity.Follow{}
	_ = f.FindByAttr(follow, attr)

	return follow.ID != 0
}

func (f *FollowRepository) FindFollow(follow *entity.Follow, attr interface{}) error {
	return f.FindByAttr(follow, attr)
}

func (u *FollowRepository) FindFollowing(username string) ([]entity.User, error) {
	var following []entity.User

	err := u.DB.Table("users").
		Preload("Following").
		Joins("INNER JOIN follows f ON users.id = f.following_id").
		Joins("INNER JOIN users follower ON f.follower_id = follower.id").
		Where("follower.username = ?", username).
		Where("f.is_private = false").
		Find(&following).
		Error
	if err != nil {
		return nil, err
	}

	return following, nil
}

func (u *FollowRepository) FindFollowers(username string) ([]entity.User, error) {
	var followers []entity.User

	err := u.DB.Table("users").
		Preload("Followers").
		Joins("INNER JOIN follows f ON users.id = f.follower_id").
		Joins("INNER JOIN users follower ON f.following_id = follower.id").
		Where("follower.username = ?", username).
		Find(&followers).
		Error
	if err != nil {
		return nil, err
	}

	return followers, nil
}

func (u *FollowRepository) FindRequestFollowers(ID uint) ([]entity.User, error) {
	var requests []entity.User

	err := u.DB.Table("users").
		Preload("Following").
		Joins("INNER JOIN follows f ON users.id = f.follower_id").
		Joins("INNER JOIN users u ON f.following_id = u.id").
		Where("u.id = ?", ID).
		Where("f.is_private = true").
		Find(&requests).
		Error
	if err != nil {
		return nil, err
	}

	return requests, nil
}
