package entity

import "time"

type Follow struct {
	ID          uint `gorm:"primaryKey"`
	FollowerID  uint `gorm:"index"`
	FollowingID uint `gorm:"index"`
	IsAccept    bool
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
