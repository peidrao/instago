package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/peidrao/instago/src/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		secretKey := os.Getenv("SECRET_KEY")
		tokenBearer := context.GetHeader("Authorization")

		if tokenBearer == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			context.Abort()
			return
		}

		tokenString := utils.ExtractBearerToken(tokenBearer)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			context.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse claims"})
			context.Abort()
			return
		}

		username, ok := claims["username"].(string)
		if !ok {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parser usename"})
			context.Abort()
			return
		}

		context.Set("username", username)
		context.Next()
	}
}
