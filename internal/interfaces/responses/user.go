package responses

import "time"

type FollowUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type UserDetailShortResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Picture  string `json:"picture"`
}

type UserDetailResponse struct {
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	Bio       string `json:"bio"`
	Link      string `json:"link"`
	IsPrivate bool   `json:"is_private"`

	Picture   string `json:"picture"`
	Followers uint   `json:"followers"`
	Following uint   `json:"following"`

	CreatedAt time.Time `json:"created_at"`
}
