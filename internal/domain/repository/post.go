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

func (p *PostRepository) CreatePost(post *entity.Post) error {
	return p.Create(&post)
}

func (p *PostRepository) FindPostsByUser(ID uint) ([]entity.Post, error) {
	var posts []entity.Post

	attr := map[string]interface{}{"user_id": ID}
	err := p.FindByAttr(&posts, attr)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostRepository) FindPostByID(ID uint) (*entity.Post, error) {
	var post entity.Post

	err := p.GetByID(&post, ID)

	if err != nil {
		return nil, err
	}

	return &post, nil
}
