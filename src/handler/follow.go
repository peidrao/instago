package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/src/domain/models"
)

func (h *UserHandler) FollowUser(context *gin.Context) {
	var request models.FolloweUserRequest

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

	followers, err := h.userRepo.FindFollowers(username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve followers"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, followers)

}

func (h *UserHandler) GetFollowings(context *gin.Context) {
	username := context.Param("username")

	followers, err := h.userRepo.FindFollowings(username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve followings"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, followers)

}
