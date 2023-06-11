package requests

type CredentialsRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type FolloweUserRequest struct {
	UserID   uint `json:"user_id"`
	FollowID uint `json:"follow_id"`
}
