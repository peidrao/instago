package permissions

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/src/domain/models"
)

func IsUserAdmin(context *gin.Context) bool {
	user, exists := context.Get("user")
	if exists {
		if u, ok := user.(*models.User); ok {
			return u.IsAdmin
		}
	}

	return false
}
