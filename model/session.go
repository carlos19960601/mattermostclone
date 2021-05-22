package model

const (
	SESSION_COOKIE_TOKEN = "MMAUTHTOKEN"
	SESSION_COOKIE_USER  = "MMUSERID"
)

type Session struct {
	Id     string `json:"id"`
	Token  string `json:"token"`
	UserId string `json:"user_id"`
}
