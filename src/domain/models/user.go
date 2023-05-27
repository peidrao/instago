package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `json:"username" gorm:"unique"`
	Email          string `json:"email" gorm:"unique"`
	Password       string `json:"password"`
	FullName       string `json:"full_name"`
	Bio            string `json:"bio"`
	Active         bool   `json:"active"`
	Private        bool   `json:"private"`
	Link           string `json:"link"`
	ProfilePicture string `json:"profile_picture"`
	// Posts          []Post
	Followers []User    `gorm:"many2many:user_followers;joinForeignKey:follower_id;joinReferences:user_id"`
	Following []User    `gorm:"many2many:user_following;joinForeignKey:user_id;joinReferences:follower_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
