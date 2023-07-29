package tests

import (
	"testing"

	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var userRepo *repository.UserRepository
var postRepo *repository.PostRepository
var followRepo *repository.FollowRepository

func setupTest() {
	var err error
	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Error connect database: " + err.Error())
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Follow{}, &entity.Post{})

	if err != nil {
		panic("Error migrating database schema: " + err.Error())
	}

	userRepo = repository.NewUserRepository(db)
	postRepo = repository.NewPostRepository(db)
	followRepo = repository.NewFollowRepository(db)
}

func tearDownTest() {
	db.Migrator().DropTable(&entity.User{}, &entity.Post{}, &entity.Follow{})
	db = nil
	userRepo = nil
}

func createUser(t *testing.T, username, email, fullName string) *entity.User {
	user := &entity.User{
		Username: username,
		Email:    email,
		Password: utils.GenerateRandomString(10),
		FullName: fullName,
	}

	err := userRepo.CreateUser(user)
	assert.NoError(t, err, "Erro ao criar usuário")
	return user
}

func createPost(t *testing.T, userID uint) *entity.Post {

	imageName := utils.GenerateRandomString(8) + ".jpg"

	post := &entity.Post{
		Caption:  utils.GenerateRandomString(10),
		Location: utils.GenerateRandomString(10),
		ImageURL: imageName,
		UserID:   userID,
	}

	err := postRepo.CreatePost(post)
	assert.NoError(t, err, "Erro ao criar postagem")
	return post
}

func createFollow(t *testing.T, followerID, followingID uint) *entity.Follow {
	follow := &entity.Follow{
		FollowerID:  followerID,
		FollowingID: followingID,
	}

	err := followRepo.CreateFollow(follow)
	assert.NoError(t, err, "Erro ao seguir usuário")
	return follow
}
