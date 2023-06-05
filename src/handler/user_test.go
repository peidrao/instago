package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/src/domain/models"
	"github.com/peidrao/instago/src/handler"
	"github.com/peidrao/instago/src/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabaseConnction() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	db.AutoMigrate(&models.User{})
	return db
}

func TestRegisterUSer(t *testing.T) {
	router := gin.Default()
	db := setupDatabaseConnction()

	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	t.Run("Should register a new user", func(t *testing.T) {
		user := models.User{
			Username: "teste",
			Email:    "teste@test.com",
			Password: "@Teste123",
			FullName: "Test One",
		}

		router.POST("/users", userHandler.RegisterUser)

		payload, err := json.Marshal(user)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/users", bytes.NewReader(payload))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)

	})
}
