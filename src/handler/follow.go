package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/src/domain/requests"
	"github.com/peidrao/instago/src/domain/responses"
)

func (h *UserHandler) FollowUser(context *gin.Context) {
	var request requests.FolloweUserRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		context.Abort()
		return
	}

	err := h.userRepo.FollowUser(request.UserID, request.FollowID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User following"})

}

func (h *UserHandler) GetFollowers(context *gin.Context) {
	username := context.Param("username")
	var response []responses.FollowUserResponse

	followers, err := h.userRepo.FindFollowers(username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve followers"})
		context.Abort()
		return
	}

	for _, following := range followers {
		follow := responses.FollowUserResponse{
			ID:       following.ID,
			Username: following.Username,
			FullName: following.FullName,
		}
		response = append(response, follow)
	}

	context.JSON(http.StatusOK, followers)

}

func (h *UserHandler) GetFollowings(context *gin.Context) {
	username := context.Param("username")
	var response []responses.FollowUserResponse

	followings, err := h.userRepo.FindFollowings(username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve followings"})
		context.Abort()
		return
	}

	for _, following := range followings {
		follow := responses.FollowUserResponse{
			ID:       following.ID,
			Username: following.Username,
			FullName: following.FullName,
		}
		response = append(response, follow)
	}

	context.JSON(http.StatusOK, response)

}
