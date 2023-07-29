package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/serializers"
	"github.com/peidrao/instago/internal/interfaces/requests"
	"github.com/peidrao/instago/internal/interfaces/responses"
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

func (h *UserHandler) CreateUserHandler(context *gin.Context) {

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

	response := serializers.UserDetailSerializer(
		&user,
		0,
		0,
	)

	context.JSON(http.StatusCreated, response)
}

func (h *UserHandler) GetUserHandler(context *gin.Context) {
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

func (h *UserHandler) GetAllUsersHandler(context *gin.Context) {
	users, err := h.UserRepository.FindAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, users)
}

func (h *UserHandler) UserMeHandler(context *gin.Context) {
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

func (h *UserHandler) UpdateUserHandler(context *gin.Context) {
	userID := context.GetUint("userID")
	var request requests.UserUpdateRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	user, followers, following, err := h.UserRepository.FindUserByID(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.UserRepository.Update(&user, request)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := serializers.UserDetailSerializer(
		user,
		followers,
		following,
	)

	context.JSON(http.StatusOK, response)
}

func (h *UserHandler) UpdatePictureUserHandler(context *gin.Context) {
	userID := context.GetUint("userID")

	user, followers, following, err := h.UserRepository.FindUserByID(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := context.FormFile("picture")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Base(file.Filename)
	err = context.SaveUploadedFile(file, "static/picture/"+filename)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	imageURL := "/static/picture/" + filename

	data := map[string]interface{}{"profile_picture": imageURL}

	err = h.UserRepository.Update(&user, data)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := serializers.UserDetailSerializer(
		user,
		followers,
		following,
	)

	context.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetSuggestionsForUserHandler(context *gin.Context) {
	userID := context.GetUint("userID")
	var usesResponse []responses.UserDetailShortResponse

	users := h.UserRepository.FindUsersSuggestions(userID)

	for _, user := range users {
		response := responses.UserDetailShortResponse{
			ID:       user.ID,
			Username: user.Username,
			FullName: user.FullName,
			Picture:  user.ProfilePicture,
		}
		usesResponse = append(usesResponse, response)
	}

	context.JSON(http.StatusOK, usesResponse)

}
