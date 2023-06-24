package repository

import (
	"github.com/peidrao/instago/internal/domain/entity"
	"gorm.io/gorm"
)

type PostRepository struct {
	*GenericRepository
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		GenericRepository: NewRepository(db),
	}
}

func (f *PostRepository) CreatePost(post *entity.Post) error {
	return f.Create(&post)
}
