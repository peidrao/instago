package repository

import (
	"log"

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

	log.Print("USER_ID -> ", ID)

	err := f.DB.Table("posts p").
		Preload("User").
		Joins("JOIN users u ON u.id = p.user_id").
		Joins("JOIN follows f ON f.following_id = u.id").
		Where("f.follower_id = ?", ID).
		Or("p.user_id = ?", ID).
		Find(&posts).
		Error
	if err != nil {
		return nil, err
	}

	return posts, nil

}
