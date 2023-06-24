package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/requests"
)

type PostHandler struct {
	PostRepository *repository.PostRepository
	UserRepository *repository.UserRepository
}

func NewPostHandler(
	userRepository *repository.UserRepository, postRepository *repository.PostRepository) *PostHandler {
	return &PostHandler{
		PostRepository: postRepository,
		UserRepository: userRepository,
	}
}

func (p *PostHandler) CreatePost(context *gin.Context) {
	var request requests.PostRequest
	userID := context.GetUint("userID")

	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := context.FormFile("image_url")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Base(file.Filename)
	err = context.SaveUploadedFile(file, "static/posts/"+filename)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	imageURL := "/static/posts/" + filename

	post := entity.Post{
		Caption:  request.Caption,
		Location: request.Location,
		ImageURL: imageURL,
		UserID:   userID,
	}

	err = p.PostRepository.CreatePost(&post)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, post)
}
