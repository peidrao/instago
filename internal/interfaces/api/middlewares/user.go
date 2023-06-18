package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/entity"
)

func SetUserMiddleware(userRepo entity.UserInterface) gin.HandlerFunc {
	return func(context *gin.Context) {
		username, _ := context.Get("username")

		str, ok := username.(string)
		if ok {
			user, err := userRepo.FindUserByUsername(str)

			if err != nil {
				context.JSON(http.StatusNotFound, gin.H{"error": "username not found in database"})
				context.Abort()
			}
			context.Set("user", user)
			context.Next()
		}
	}
}