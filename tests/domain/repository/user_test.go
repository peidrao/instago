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

	err = db.AutoMigrate(&entity.User{})

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

func TestUserCreateRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	user := &entity.User{
		Username: "teste",
		Email:    "teste@gmail.com",
		Password: utils.GenerateRandomString(10),
		FullName: "Teste OK",
	}

	err := userRepo.CreateUser(user)
	assert.NoError(t, err, "Erro ao criar usuário")

	var savedUser entity.User
	err = db.First(&savedUser, user.ID).Error
	assert.NoError(t, err, "Erro ao buscar usuário do banco de dados")
	assert.Equal(t, user.Username, savedUser.Username, "Nomes de usuário não coincidem")
	assert.Equal(t, user.Email, savedUser.Email, "E-mails não coincidem")
	assert.Equal(t, user.Password, savedUser.Password, "Senhas não coincidem")
}

func TestFindAllUsersRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	users := []entity.User{
		{Username: "user1", Email: "user1@example.com", Password: utils.GenerateRandomString(5)},
		{Username: "user2", Email: "user2@example.com", Password: utils.GenerateRandomString(5)},
		{Username: "user3", Email: "user3@example.com", Password: utils.GenerateRandomString(5)},
		{Username: "user4", Email: "user4@example.com", Password: utils.GenerateRandomString(5)},
		{Username: "user5", Email: "user5@example.com", Password: utils.GenerateRandomString(5)},
	}

	for i := range users {
		if err := userRepo.CreateUser(&users[i]); err != nil {
			t.Fatalf("Erro ao criar usuário: %v", err)
		}
	}

	resultUsers, err := userRepo.FindAllUsers()
	assert.NoError(t, err, "Erro ao buscar usuários")
	assert.Len(t, resultUsers, len(users), "Quantidade de usuários está incorreta")

	for i, user := range users {
		assert.Equal(t, user.Username, resultUsers[i].Username, "Nomes de usários não coincidem")
		assert.Equal(t, user.Email, resultUsers[i].Email, "E-mails de usários não coincidem")
		assert.Equal(t, user.Password, resultUsers[i].Password, "Senhas de usários não coincidem")
	}
}

