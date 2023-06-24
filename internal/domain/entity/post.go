package entity

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"column:user_id"`    // Chave estrangeira para o ID do usuário
	User      User `gorm:"foreignKey:UserID"` // Relacionamento com o usuário
	ImageURL  string
	Caption   string
	Location  string
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
