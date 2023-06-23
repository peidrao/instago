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

// func (u *UserRepository) FindFollowers(username string) ([]*entity.User, error) {
// 	var user entity.User
// 	result := u.db.Preload("Followers").Preload("Following").Where("username = ?", username).First(&user)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return user.Followers, nil
// }

// func (u *UserRepository) FindFollowing(username string) ([]*entity.User, error) {
// 	var user entity.User
// 	result := u.db.Preload("Followers").Preload("Following").Where("username = ?", username).First(&user)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return user.Following, nil
// }
