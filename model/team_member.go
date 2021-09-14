package model

type TeamMember struct {
	TeamId   string `json:"team_id"`
	UserId   string `json:"user_id"`
	Roles    string `json:"roles"`
	DeleteAt int64  `json:"delete_at"`
}
