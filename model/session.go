package model

const (
	SESSION_COOKIE_TOKEN = "MMAUTHTOKEN"
	SESSION_COOKIE_USER  = "MMUSERID"
)

type Session struct {
	Id          string        `json:"id"`
	Token       string        `json:"token"`
	UserId      string        `json:"user_id"`
	TeamMembers []*TeamMember `json:"team_members" db:"-"`
}

func (s *Session) PreSave() {
	if s.Id == "" {
		s.Id = NewId()
	}
	if s.Token == "" {
		s.Token = NewId()
	}
}
