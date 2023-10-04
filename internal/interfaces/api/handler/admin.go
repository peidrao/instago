package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/responses"
	"net/http"
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

	usersResponse := make([]responses.UserAllDetailResponse, len(users))

	for i, user := range users {
		usersResponse[i] = responses.UserAllDetailResponse{
			ID:             user.ID,
			Username:       user.Username,
			FullName:       user.FullName,
			Email:          user.Email,
			Password:       user.Password,
			Bio:            user.Bio,
			Link:           user.Link,
			ProfilePicture: user.ProfilePicture,
			Active:         user.IsActive,
			Private:        user.IsPrivate,
			Admin:          user.IsAdmin,
			CreatedAt:      user.CreatedAt,
			UpdatedAt:      user.UpdatedAt,
		}
	}

	context.JSON(http.StatusOK, usersResponse)
}
