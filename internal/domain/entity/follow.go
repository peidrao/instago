package entity

import "time"

type Follow struct {
	ID          uint      `gorm:"primaryKey"`
	FollowerID  uint      `gorm:"index"`
	FollowingID uint      `gorm:"index"`
	IsPrivate   bool      `gorm:"default:false"`
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
