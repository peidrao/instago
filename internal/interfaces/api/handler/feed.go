package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/api/serializers"
	"github.com/peidrao/instago/internal/interfaces/responses"
	"github.com/peidrao/instago/utils"
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

func (f *FeedHandler) FeedMeHandler(context *gin.Context) {
	var postResponse []responses.PostDetailResponse
	page, pageSize := utils.ParserPageAndPageSize(context)

	userID := context.GetUint("userID")

	feed, err := f.FeedRepository.GetMeFeed(userID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	paginatedFeed := utils.GenericPaginated(feed, page, pageSize).([]entity.Post)

	totalItems := len(feed)
	for _, post := range paginatedFeed {
		serializedPost := serializers.PostDetailSerializer(&post)
		postResponse = append(postResponse, *serializedPost)
	}

	nextLink, prevLink := utils.CalculateLinks(context, page, pageSize, totalItems)

	response := gin.H{
		"count":    totalItems,
		"next":     nextLink,
		"previous": prevLink,
		"results":  postResponse,
	}

	context.JSON(http.StatusOK, response)
}
