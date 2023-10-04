package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
)

type AdminHandler struct {
	AdminRepository *repository.AdminRepository
}

func NewAdminHandler(adminRepository *repository.AdminRepository) *AdminHandler {
	return &AdminHandler{
		AdminRepository: adminRepository,
	}
}

func (a *AdminHandler) GetAllUsersHandler(context *gin.Context) {
	users, err := a.AdminRepository.FindAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, users)
}
