package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/responses"
)

type FeedHandler struct {
	PostRepository *repository.PostRepository
	UserRepository *repository.UserRepository
	FeedRepository *repository.FeedRepository
}

func NewFeedHandler(
	userRepository *repository.UserRepository,
	postRepository *repository.PostRepository,
	feedRepository *repository.FeedRepository,
) *FeedHandler {
	return &FeedHandler{
		PostRepository: postRepository,
		UserRepository: userRepository,
		FeedRepository: feedRepository,
	}
}

func (f *FeedHandler) FeedMe(context *gin.Context) {
	var postResponse []responses.PostDetailResponse

	userID := context.GetUint("userID")

	feed, err := f.FeedRepository.GetMeFeed(userID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	for _, f := range feed {
		response := responses.PostDetailResponse{
			ID:        f.ID,
			UserID:    f.UserID,
			ImageURL:  f.ImageURL,
			Caption:   f.Caption,
			Location:  f.Location,
			CreatedAt: f.CreatedAt,
		}
		postResponse = append(postResponse, response)
	}

	context.JSON(http.StatusOK, postResponse)
}
