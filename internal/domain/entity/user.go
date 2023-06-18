package entity

import (
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

	ProfilePicture string `json:"profile_picture" gorm:"null"`
	// Posts          []Post

	Followers []*User `gorm:"many2many:user_followers;joinForeignKey:follower_id;joinReferences:followed_id"`
	Following []*User `gorm:"many2many:user_followers;joinForeignKey:followed_id;joinReferences:follower_id"`

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
	validate.RegisterValidation("strongPassword", StrongPasswordValidator)
	return validate
}

type UserInterface interface {
	CreateUser(user *User) error
	FindUserByID(id uint) (*User, error)
	FindUserByEmail(email string) (*User, error)
	FindUserByUsername(username string) (*User, error)
	FindLastUser() (*User, error)
	ListAllUsers() ([]*User, error)
	DestroyUser(id uint64) error
	UpdateUser(*User) (*User, error)

	FollowUser(userId, followerID uint) error
	UnFollowUser(userId, unfollowerID uint) error
	FindFollowers(username string) ([]*User, error)
	FindFollowing(username string) ([]*User, error)
}