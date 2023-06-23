package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/requests"
)

type FollowHandler struct {
	FollowRepository *repository.FollowRepository
	UserRepository   *repository.UserRepository
}

func NewFollowHandler(
	userRepository *repository.UserRepository, followRepository *repository.FollowRepository) *FollowHandler {
	return &FollowHandler{
		FollowRepository: followRepository,
		UserRepository:   userRepository,
	}
}

func (f *FollowHandler) FollowUser(context *gin.Context) {
	var request requests.FolloweUserRequest
	var newFollow entity.Follow
	user, _ := context.Get("user")

	userObj, _ := user.(*entity.User)
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	followUser, err := f.UserRepository.FindUserByID(request.FollowID)
	log.Println("NEW FOLLOW -> ", newFollow)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if followUser.IsPrivate {
		newFollow.IsPrivate = true
	}

	newFollow.FollowerID = userObj.ID
	newFollow.FollowingID = followUser.ID

	log.Println("NEW FOLLOW -> ", newFollow)

	err = f.FollowRepository.CreateFollow(&newFollow)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User following"})
}

func (f *FollowHandler) UnfollowUser(context *gin.Context) {
	var request requests.FolloweUserRequest
	var follow entity.Follow
	user, _ := context.Get("user")

	userObj, _ := user.(*entity.User)

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		context.Abort()
		return
	}
	follow.FollowingID = request.FollowID
	follow.FollowerID = userObj.ID

	err := f.FollowRepository.FindByAttr(follow, &entity.Follow{})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	attr := map[string]interface{}{"is_active": false}

	err = f.FollowRepository.UpdateFollow(follow, attr)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Follower removal successful."})

}

// func (h *UserHandler) GetFollowers(context *gin.Context) {
// 	username := context.Param("username")
// 	var response []responses.FollowUserResponse

// 	followers, err := h.userRepo.FindFollowers(username)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve followers"})
// 		context.Abort()
// 		return
// 	}

// 	for _, following := range followers {
// 		follow := responses.FollowUserResponse{
// 			ID:       following.ID,
// 			Username: following.Username,
// 			FullName: following.FullName,
// 		}
// 		response = append(response, follow)
// 	}

// 	if len(followers) == 0 {
// 		response = make([]responses.FollowUserResponse, 0)
// 	}

// 	context.JSON(http.StatusOK, response)
// }

// func (h *UserHandler) GetFollowing(context *gin.Context) {
// 	username := context.Param("username")
// 	var response []responses.FollowUserResponse

// 	following, err := h.userRepo.FindFollowing(username)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve following"})
// 		context.Abort()
// 		return
// 	}

// 	for _, following := range following {
// 		follow := responses.FollowUserResponse{
// 			ID:       following.ID,
// 			Username: following.Username,
// 			FullName: following.FullName,
// 		}
// 		response = append(response, follow)
// 	}

// 	context.JSON(http.StatusOK, response)
// }
