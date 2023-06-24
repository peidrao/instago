package requests

type CredentialsRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type FolloweUserRequest struct {
	FollowID uint `json:"follow_id"`
}

type UserIDRequest struct {
	ID uint `json:"id"`
}
