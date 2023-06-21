package handler

// func (h *UserHandler) FollowUser(context *gin.Context) {
// 	var request requests.FolloweUserRequest
// 	user, _ := context.Get("user")

// 	userObj, _ := user.(*entity.User)

// 	if err := context.ShouldBindJSON(&request); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		context.Abort()
// 		return
// 	}

// 	err := h.userRepo.FollowUser(userObj.ID, request.FollowID)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user"})
// 		context.Abort()
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"message": "User following"})
// }

// func (h *UserHandler) UnfollowUser(context *gin.Context) {
// 	var request requests.FolloweUserRequest
// 	user, _ := context.Get("user")

// 	userObj, _ := user.(*entity.User)

// 	if err := context.ShouldBindJSON(&request); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		context.Abort()
// 		return
// 	}

// 	err := h.userRepo.UnFollowUser(userObj.ID, request.FollowID)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		context.Abort()
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"message": "Follower removal successful."})

// }

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
