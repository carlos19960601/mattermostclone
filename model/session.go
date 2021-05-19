package model

type Session struct {
	Id     string `json:"id"`
	Token  string `json:"token"`
	UserId string `json:"user_id"`
}
