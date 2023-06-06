package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/peidrao/instago/src/domain/models"
	"github.com/peidrao/instago/src/repository"

	"github.com/peidrao/instago/src/utils"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (h *UserHandler) RegisterUser(context *gin.Context) {

	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	validate := models.NewValidator()

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
	}

	user.Password, _ = utils.HashPassword(user.Password)
	user.Active = true

	err = h.userRepo.CreateUser(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) GetUser(context *gin.Context) {
	id := context.Param("id")

	var userID uint

	if _, err := fmt.Sscan(id, &userID); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	log.Print("\n\n\nOI")
	log.Print(userID)
	log.Print("\n\n\n")

	user, err := h.userRepo.GetUserByID(userID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (h *UserHandler) RemoveUser(context *gin.Context) {
	id := context.Param("id")

	userID, _ := strconv.ParseUint(id, 10, 64)

	err := h.userRepo.RemoveUser(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
