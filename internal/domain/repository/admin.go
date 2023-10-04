package repository

import (
	"github.com/peidrao/instago/internal/domain/entity"
	"gorm.io/gorm"
)

type AdminRepository struct {
	*GenericRepository
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{
		GenericRepository: NewRepository(db),
	}
}

func (a *AdminRepository) FindAllUsers() ([]entity.User, error) {
	var users []entity.User
	if err := a.FindAll(&users); err != nil {
		return nil, err
	}

	return users, nil
}
