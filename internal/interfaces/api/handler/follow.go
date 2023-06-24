package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/peidrao/instago/internal/domain/repository"
	"github.com/peidrao/instago/internal/interfaces/requests"
	"github.com/peidrao/instago/internal/interfaces/responses"
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
	var follow entity.Follow
	user, _ := context.Get("user")

	userObj, _ := user.(*entity.User)
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	userFollowing, _, _, err := f.UserRepository.FindUserByID(request.FollowID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	query := map[string]interface{}{"follower_id": userObj.ID, "following_id": userFollowing.ID}
	exists := f.FollowRepository.FindLinkFollows(query)

	if exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "AI CALICA"})
		context.Abort()
		return
	}

	if !userFollowing.IsPrivate {
		follow.IsAccept = true
	}

	follow.FollowerID = userObj.ID
	follow.FollowingID = userFollowing.ID

	err = f.FollowRepository.CreateFollow(&follow)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user"})
		context.Abort()
		return
	}

	log.Println("FOLLOW ACCEPT -> ", follow.IsAccept)

	if userFollowing.IsPrivate {
		context.JSON(http.StatusOK, gin.H{"message": "Request sent!"})
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
	query := map[string]interface{}{"following_id": request.FollowID, "follower_id": userObj.ID}

	err := f.FollowRepository.FindFollow(&follow, query)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	attr := map[string]interface{}{"is_active": false}

	err = f.FollowRepository.UpdateFollow(&follow, attr)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Follower removal successful."})
}

func (f *FollowHandler) GetFollowing(context *gin.Context) {
	username := context.Param("username")
	var response []responses.FollowUserResponse

	followers, err := f.FollowRepository.FindFollowing(username)

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

	if len(followers) == 0 {
		response = make([]responses.FollowUserResponse, 0)
	}

	context.JSON(http.StatusOK, response)
}

func (f *FollowHandler) GetFollowers(context *gin.Context) {
	username := context.Param("username")
	var response []responses.FollowUserResponse

	followers, err := f.FollowRepository.FindFollowers(username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve followers"})
		context.Abort()
		return
	}

	for _, follower := range followers {
		follow := responses.FollowUserResponse{
			ID:       follower.ID,
			Username: follower.Username,
			FullName: follower.FullName,
		}
		response = append(response, follow)
	}

	if len(followers) == 0 {
		response = make([]responses.FollowUserResponse, 0)
	}

	context.JSON(http.StatusOK, response)
}

func (f *FollowHandler) GetFollowersRequest(context *gin.Context) {
	var response []responses.FollowUserResponse

	user, _ := context.Get("userID")

	userID, _ := user.(uint)

	followers, err := f.FollowRepository.FindRequestFollowers(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	context.JSON(http.StatusOK, response)
}

func (f *FollowHandler) AcceptRequest(context *gin.Context) {
	var request requests.UserIDRequest
	var follow entity.Follow

	user, _ := context.Get("userID")

	userID, _ := user.(uint)

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		context.Abort()
		return
	}

	attr := map[string]interface{}{"follower_id": request.ID, "following_id": userID}

	err := f.FollowRepository.FindFollow(&follow, attr)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	attr = map[string]interface{}{"is_accept": true}

	err = f.FollowRepository.UpdateFollow(&follow, attr)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Ok"})
}

func (f *FollowHandler) GetFollowingRequest(context *gin.Context) {
	var response []responses.FollowUserResponse

	userID := context.GetUint("userID")

	followers, err := f.FollowRepository.FindRequestFollowing(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	context.JSON(http.StatusOK, response)
}

func (f *FollowHandler) CancelRequest(context *gin.Context) {
	var request requests.UserIDRequest
	var follow entity.Follow

	user, _ := context.Get("userID")

	userID, _ := user.(uint)

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		context.Abort()
		return
	}

	attr := map[string]interface{}{"follower_id": userID, "following_id": request.ID}

	err := f.FollowRepository.FindFollow(&follow, attr)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	err = f.FollowRepository.DeleteFollow(&follow, follow.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Delete"})
}
