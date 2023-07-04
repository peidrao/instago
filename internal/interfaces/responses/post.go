package responses

import "time"

type PostDetailResponse struct {
	ID        uint                    `json:"id"`
	ImageURL  string                  `json:"image_url"`
	Caption   string                  `json:"caption"`
	Location  string                  `json:"location"`
	CreatedAt time.Time               `json:"created_at"`
	User      UserDetailShortResponse `json:"user"`
}
