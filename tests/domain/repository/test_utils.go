package repository_test

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

func setupTest() {
	var err error
	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Error connect database: " + err.Error())
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Follow{})

	if err != nil {
		panic("Error migrating database schema: " + err.Error())
	}

	userRepo = repository.NewUserRepository(db)
}

func tearDownTest() {
	db.Migrator().DropTable(&entity.User{})
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
