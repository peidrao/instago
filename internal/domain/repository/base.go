package repository

import (
	"gorm.io/gorm"
)

type GenericRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *GenericRepository {
	return &GenericRepository{
		DB: db,
	}
}

func (r *GenericRepository) Create(data interface{}) error {
	if err := r.DB.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (r *GenericRepository) GetByID(id uint) (interface{}, error) {
	var result interface{}

	if err := r.DB.First(&result, id).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *GenericRepository) FindAll(entities interface{}) error {
	if err := r.DB.Find(entities).Error; err != nil {
		return err
	}

	return nil
}

func (r *GenericRepository) FindByAttr(attr string, value interface{}, entity interface{}) error {
	if err := r.DB.Where(attr+" = ?", value).First(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *GenericRepository) Update(data interface{}, id uint) error {
	err := r.DB.Model(data).Where("id = ?", id).Updates(data).Error
	return err
}

func (r *GenericRepository) Delete(id uint) error {
	err := r.DB.Delete(nil, id).Error
	return err
}
