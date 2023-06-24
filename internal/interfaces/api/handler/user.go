package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/serializers"
	"github.com/peidrao/instago/utils"
)

type UserHandler struct {
	UserRepository *repository.UserRepository
}

func NewUserHandler(userRepository *repository.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
	}
}

func (h *UserHandler) CreateUser(context *gin.Context) {

	var user entity.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	validate := entity.NewValidator()

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
	}

	_, errFindEmail := h.UserRepository.FindUserByEmail(user.Email)

	if errFindEmail != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "This email already exists for a user"})
		context.Abort()
		return
	}

	user.Password, _ = utils.HashPassword(user.Password)

	err = h.UserRepository.CreateUser(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) GetUser(context *gin.Context) {
	username := context.Param("username")

	user, followers, following, err := h.UserRepository.FindUserByUsername(username)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	response := serializers.UserDetailSerializer(
		user,
		followers,
		following,
	)

	context.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetAllUsers(context *gin.Context) {
	users, err := h.UserRepository.FindAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, users)
}

func (h *UserHandler) UserMe(context *gin.Context) {
	userID := context.GetUint("userID")

	user, followers, following, err := h.UserRepository.FindUserByID(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	response := serializers.UserDetailSerializer(
		user,
		followers,
		following,
	)

	context.JSON(http.StatusOK, response)
}
