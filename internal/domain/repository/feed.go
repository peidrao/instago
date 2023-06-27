package repository

import (
	"github.com/peidrao/instago/internal/domain/entity"
	"gorm.io/gorm"
)

type FeedRepository struct {
	*GenericRepository
}

func NewFeedRepository(db *gorm.DB) *FeedRepository {
	return &FeedRepository{
		GenericRepository: NewRepository(db),
	}
}

func (f *FeedRepository) GetMeFeed(ID uint) ([]entity.Post, error) {
	var posts []entity.Post

	err := f.DB.Table("posts p").
		Preload("User").
		Joins("JOIN users u ON u.id = p.user_id").
		Joins("JOIN follows f ON f.following_id = u.id").
		Where("f.follower_id = ?", ID).
		Where("f.is_active = true").
		Or("p.user_id = ?", ID).
		Find(&posts).
		Error
	if err != nil {
		return nil, err
	}

	return posts, nil

}
