package requests

type PostRequest struct {
	Caption  string `form:"caption"`
	Location string `form:"location"`
}
