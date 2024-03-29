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
	return f.Create(&follow)
}

func (f *FollowRepository) DeleteFollow(follow *entity.Follow, ID uint) error {
	return f.Delete(&follow, ID)
}

func (f *FollowRepository) UpdateFollow(follow *entity.Follow, attr interface{}) error {
	return f.Update(&follow, attr)
}

func (f *FollowRepository) FindLinkFollows(attr interface{}) bool {
	var follow entity.Follow
	_ = f.FindByAttr(&follow, attr)

	return follow.ID != 0
}

func (f *FollowRepository) FindFollow(follow *entity.Follow, attr interface{}) error {
	return f.FindByAttr(&follow, attr)
}

func (f *FollowRepository) FindFollowing(username string) ([]entity.User, error) {
	var following []entity.User

	err := f.DB.Table("users").
		Preload("Following").
		Joins("INNER JOIN follows f ON users.id = f.following_id").
		Joins("INNER JOIN users follower ON f.follower_id = follower.id").
		Where("follower.username = ?", username).
		Where("f.is_accept = true").
		Where("f.is_active = true").
		Find(&following).
		Error
	if err != nil {
		return nil, err
	}

	return following, nil
}

func (f *FollowRepository) FindFollowers(username string) ([]entity.User, error) {
	var followers []entity.User

	err := f.DB.Table("users").
		Preload("Followers").
		Joins("INNER JOIN follows f ON users.id = f.follower_id").
		Joins("INNER JOIN users u ON f.following_id = u.id").
		Where("u.username = ?", username).
		Where("f.is_accept = true").
		Where("f.is_active = true").
		Find(&followers).
		Error
	if err != nil {
		return nil, err
	}

	return followers, nil
}

func (f *FollowRepository) FindRequestFollowers(ID uint) ([]entity.User, error) {
	var requests []entity.User

	err := f.DB.Table("users").
		Preload("Following").
		Joins("INNER JOIN follows f ON users.id = f.follower_id").
		Joins("INNER JOIN users u ON f.following_id = u.id").
		Where("u.id = ?", ID).
		Where("f.is_accept = false").
		Find(&requests).
		Error
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (f *FollowRepository) FindRequestFollowing(ID uint) ([]entity.User, error) {
	var requests []entity.User

	err := f.DB.Table("users").
		Preload("Following").
		Joins("INNER JOIN follows f ON users.id = f.following_id").
		Joins("INNER JOIN users u ON f.follower_id = u.id").
		Where("u.id = ?", ID).
		Where("f.is_accept = false").
		Find(&requests).
		Error
	if err != nil {
		return nil, err
	}

	return requests, nil
}
