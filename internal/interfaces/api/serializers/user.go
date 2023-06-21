package serializers

import (
	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/peidrao/instago/internal/interfaces/responses"
)

func UserDetailSerializer(user *entity.User) *responses.UserDetailResponse {

	return &responses.UserDetailResponse{
		Username: user.Username,
		FullName: user.FullName,
		Bio:      user.Bio,
		Link:     user.Link,

		Picture: user.ProfilePicture,

		CreatedAt: user.CreatedAt,
	}
}
