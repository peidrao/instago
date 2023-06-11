package permissions

import (
	"github.com/gin-gonic/gin"
	"github.com/peidrao/instago/internal/domain/entity"
)

func IsUserAdminPermission(context *gin.Context) bool {
	user, exists := context.Get("user")
	if exists {
		if u, ok := user.(*entity.User); ok {
			return u.IsAdmin
		}
	}

	return false
}
