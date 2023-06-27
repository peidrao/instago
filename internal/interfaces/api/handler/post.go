package handler

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/requests"
	"github.com/peidrao/instago/internal/interfaces/responses"
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

func (p *PostHandler) GetMePosts(context *gin.Context) {
	var postResponse []responses.PostDetailResponse
	userID := context.GetUint("userID")

	posts, err := p.PostRepository.FindPostsByUser(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, post := range posts {
		response := responses.PostDetailResponse{
			ID:        post.ID,
			UserID:    post.UserID,
			ImageURL:  post.ImageURL,
			Caption:   post.Caption,
			Location:  post.Location,
			CreatedAt: post.CreatedAt,
		}
		postResponse = append(postResponse, response)
	}

	context.JSON(http.StatusOK, postResponse)
}

func (p *PostHandler) GetPost(context *gin.Context) {
	ID := context.Param("id")

	uintID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	post, err := p.PostRepository.FindPostByID(uint(uintID))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := responses.PostDetailResponse{
		ID:        post.ID,
		UserID:    post.UserID,
		ImageURL:  post.ImageURL,
		Caption:   post.Caption,
		Location:  post.Location,
		CreatedAt: post.CreatedAt,
	}

	context.JSON(http.StatusOK, response)
}

func (p *PostHandler) DeletePost(context *gin.Context) {
	ID := context.Param("id")

	uintID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = p.PostRepository.RemovePost(uint(uintID))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Post removed"})
}
