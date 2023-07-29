package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/interfaces/requests"
	"github.com/peidrao/instago/internal/interfaces/responses"
	"github.com/peidrao/instago/utils"
)

func (h *UserHandler) LoginHandler(context *gin.Context) {
	var credentials requests.CredentialsRequest

	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _, _, err := h.UserRepository.FindUserByUsername(credentials.Username)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := utils.ComparePassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	response := responses.TokenResponse{Token: token}
	context.JSON(http.StatusOK, response)
}

func (h *UserHandler) TokenIsValidHandler(context *gin.Context) {
	var tokenRequest requests.TokenRequest

	if err := context.ShouldBindJSON(&tokenRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.VerifyToken(tokenRequest.Token)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"valid": token.Valid})
}
