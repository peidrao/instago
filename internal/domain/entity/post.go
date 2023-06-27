package entity

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"column:user_id"`
	User      User `gorm:"foreignKey:UserID"`
	ImageURL  string
	Caption   string
	Location  string
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
