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

func TestUserCreateRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	user := createUser(t, "test", "test@test.com", "Test")

	var savedUser entity.User
	err := db.First(&savedUser, user.ID).Error

	assert.NoError(t, err, "Erro ao buscar usuário do banco de dados")
	assert.Equal(t, user.Username, savedUser.Username, "Nomes de usuário não coincidem")
	assert.Equal(t, user.Email, savedUser.Email, "E-mails não coincidem")
	assert.Equal(t, user.Password, savedUser.Password, "Senhas não coincidem")
}

func TestFindAllUsersRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	users := []entity.User{
		{Username: "user1", Email: "user1@example.com"},
		{Username: "user2", Email: "user2@example.com"},
		{Username: "user3", Email: "user3@example.com"},
		{Username: "user4", Email: "user4@example.com"},
		{Username: "user5", Email: "user5@example.com"},
	}

	for _, u := range users {
		createUser(t, u.Username, u.Email, u.FullName)
	}

	resultUsers, err := userRepo.FindAllUsers()
	assert.NoError(t, err, "Erro ao buscar usuários")
	assert.Len(t, resultUsers, len(users), "Quantidade de usuários está incorreta")

	for i, user := range users {
		assert.Equal(t, user.Username, resultUsers[i].Username, "Nomes de usários não coincidem")
		assert.Equal(t, user.Email, resultUsers[i].Email, "E-mails de usários não coincidem")
	}
}

func TestFindUsersByUsernameRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	user := createUser(t, "test", "test@gmail.com", "Test")

	userByUsernamed, _, _, err := userRepo.FindUserByUsername("test")
	assert.NoError(t, err, "Erro ao buscar usuários")

	assert.Equal(t, user.Username, userByUsernamed.Username, "Nomes de usários não coincidem")
	assert.Equal(t, user.Email, userByUsernamed.Email, "E-mails de usários não coincidem")
}

func TestFindUsersByIDRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	user := createUser(t, "test", "test@gmail.com", "Test")

	userByUsernamed, _, _, err := userRepo.FindUserByID(user.ID)
	assert.NoError(t, err, "Erro ao buscar usuários")

	assert.Equal(t, user.Username, userByUsernamed.Username, "Nomes de usários não coincidem")
	assert.Equal(t, user.Email, userByUsernamed.Email, "E-mails de usários não coincidem")
	assert.Equal(t, user.Password, userByUsernamed.Password, "Senhas de usários não coincidem")
}

func TestFindUsersByEmailRepository(t *testing.T) {
	setupTest()
	defer tearDownTest()

	user := createUser(t, "test", "test@gmail.com", "Test")

	userByUsernamed, err := userRepo.FindUserByEmail(user.Email)
	assert.NoError(t, err, "Erro ao buscar usuários")

	assert.Equal(t, user.Username, userByUsernamed.Username, "Nomes de usários não coincidem")
	assert.Equal(t, user.Email, userByUsernamed.Email, "E-mails de usários não coincidem")
}
