package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/repository"
)

func SetUserMiddleware(userRepository *repository.UserRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		username, _ := context.Get("username")
		str, ok := username.(string)
		if ok {
			user, _, _, err := userRepository.FindUserByUsername(str)

			if err != nil {
				context.JSON(http.StatusNotFound, gin.H{"error": "username not found in database"})
				context.Abort()
			}
			context.Set("user", user)
			context.Set("userID", user.ID)
			context.Next()
		}
	}
}
