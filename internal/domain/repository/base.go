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

func (r *GenericRepository) GetByID(result interface{}, id uint) error {
	if err := r.DB.First(&result, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *GenericRepository) FindAll(entities interface{}) error {
	if err := r.DB.Find(entities).Error; err != nil {
		return err
	}

	return nil
}

func (r *GenericRepository) FindByAttr(entity interface{}, attr interface{}) error {
	if err := r.DB.Where(attr).Find(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *GenericRepository) Update(entity interface{}, attr interface{}) error {
	err := r.DB.Model(entity).Updates(attr).Error
	return err
}

func (r *GenericRepository) Delete(entity interface{}, ID uint) error {
	err := r.DB.Delete(entity, ID).Error
	return err
}
