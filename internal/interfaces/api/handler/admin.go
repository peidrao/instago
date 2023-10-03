package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
)

type AdminHandler struct {
	UserRepository *repository.UserRepository
}

func NewAdminHandler(userRepository *repository.UserRepository) *AdminHandler {
	return &AdminHandler{
		UserRepository: userRepository,
	}
}

func (a *AdminHandler) GetAllUsersHandler(context *gin.Context) {
	users, err := a.UserRepository.FindAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, users)
}
