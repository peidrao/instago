package entity

import (
	"log"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password" validate:"required,strongPassword"`
	FullName  string `json:"full_name"`
	Bio       string `json:"bio" gorm:"null"`
	IsActive  bool   `json:"active" gorm:"default:true"`
	IsPrivate bool   `json:"private" gorm:"default:false"`
	IsAdmin   bool   `json:"admin" gorm:"default:false"`
	Link      string `json:"link" gorm:"null"`

	ProfilePicture string   `json:"profile_picture" gorm:"null"`
	Posts          []Post   `json:"posts" gorm:"foreignKey:UserID"`
	Followers      []Follow `gorm:"foreignKey:FollowerID"`
	Following      []Follow `gorm:"foreignKey:FollowingID"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func StrongPasswordValidator(f1 validator.FieldLevel) bool {
	password := f1.Field().String()
	rules := []string{
		"[a-z]",        // pelo menos uma letra minuscula
		"[A-Z]",        // pelo menos uma letra maiscula
		"[0-9]",        // pelo menos um numero
		"[!@#$%^&*()]", // pelo menos um catactere epescial
		".{8,}",        // minimo de 8 caracteres
	}

	for _, rule := range rules {
		matched, _ := regexp.MatchString(rule, password)
		if !matched {
			return false
		}
	}

	return true
}

func NewValidator() *validator.Validate {
	validate := validator.New()
	if err := validate.RegisterValidation("strongPassword", StrongPasswordValidator); err != nil {
		log.Fatalf("error validate: %v", err)
	}
	return validate
}

type UserInterface interface {
	FollowUser(userId, followerID uint) error
	UnFollowUser(userId, unfollowerID uint) error
	FindFollowers(username string) ([]*User, error)
	FindFollowing(username string) ([]*User, error)
}
