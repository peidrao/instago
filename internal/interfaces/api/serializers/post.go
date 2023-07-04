package serializers

import (
	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/peidrao/instago/internal/interfaces/responses"
)

func PostDetailSerializer(post *entity.Post) *responses.PostDetailResponse {
	user := post.User

	userResponse := responses.UserDetailShortResponse{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Picture:  user.ProfilePicture,
	}

	return &responses.PostDetailResponse{
		ID:        post.ID,
		ImageURL:  post.ImageURL,
		Caption:   post.Caption,
		Location:  post.Location,
		CreatedAt: post.CreatedAt,
		User:      userResponse,
	}
}
