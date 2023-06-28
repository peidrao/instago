package requests

type CredentialsRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	FullName  string `json:"full_name"`
	Bio       string `json:"bio"`
	IsPrivate bool   `json:"private"`
	Link      string `json:"link"`
}

type FolloweUserRequest struct {
	FollowID uint `json:"follow_id"`
}

type UserIDRequest struct {
	ID uint `json:"id"`
}
