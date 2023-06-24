package serializers

import (
	"github.com/peidrao/instago/internal/domain/entity"
	"github.com/peidrao/instago/internal/interfaces/responses"
)

func PostDetailSerializer(post *entity.Post) *responses.PostDetailResponse {

	return &responses.PostDetailResponse{
		ID:        post.ID,
		UserID:    post.UserID,
		ImageURL:  post.ImageURL,
		Caption:   post.Caption,
		Location:  post.Location,
		CreatedAt: post.CreatedAt,
	}
}
