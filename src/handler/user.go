package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/peidrao/instago/src/domain/interfaces"
	"github.com/peidrao/instago/src/domain/models"

	"github.com/peidrao/instago/src/utils"
)

type UserHandler struct {
	userRepo interfaces.UserInterface
}

func NewUserHandler(userRepo interfaces.UserInterface) *UserHandler {
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

	_, errFindEmail := h.userRepo.FindUserByEmail(user.Email)

	if errFindEmail == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "This email already exists for a user"})
		context.Abort()
		return
	}

	user.Password, _ = utils.HashPassword(user.Password)

	err = h.userRepo.CreateUser(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) GetUser(context *gin.Context) {
	username := context.Param("username")

	user, err := h.userRepo.FindUserByUsername(username)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAllUsers(context *gin.Context) {
	value, exist := context.Get("logged_in")

	log.Println(value, exist)

	users, err := h.userRepo.ListAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, users)
}

func (h *UserHandler) RemoveUser(context *gin.Context) {
	userContext, exists := context.Get("user")

	if !exists {
		context.JSON(http.StatusNotFound, gin.H{"error": "user not found in database"})
		context.Abort()
		return
	}

	if user, ok := userContext.(*models.User); ok {
		user.IsActive = false
		user.UpdatedAt = time.Now()

		_, err := h.userRepo.UpdateUser(user)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
