package store

import (
	"context"

	"github.com/zengqiang96/mattermostclone/model"
)

type Store interface {
	Channel() ChannelStore
	Post() PostStore
	Team() TeamStore
	User() UserStore
	Session() SessionStore
}

type ChannelStore interface {
	Get(id string) (*model.Channel, error)
}

type PostStore interface {
	Save(post *model.Post) (*model.Post, error)
}

type TeamStore interface {
	Get(id string) (*model.Team, error)
}

type UserStore interface {
	Count(options model.UserCountOptions) (int64, error)
	Get(ctx context.Context, id string) (*model.User, error)
	GetForLogin(loginId string) (*model.User, error)
	Save(user *model.User) (*model.User, error)
}

type SessionStore interface {
	Save(*model.Session) (*model.Session, error)
}
