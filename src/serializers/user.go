package serializers

import (
	"github.com/peidrao/instago/src/domain/models"
	"github.com/peidrao/instago/src/domain/responses"
)

func UserDetailSerializer(user *models.User, followers, following uint) *responses.UserDetailResponse {

	return &responses.UserDetailResponse{
		Username: user.Username,
		FullName: user.FullName,
		Bio:      user.Bio,
		Link:     user.Link,

		Picture:   user.ProfilePicture,
		Followers: followers,
		Following: following,
		CreatedAt: user.CreatedAt,
	}
}
