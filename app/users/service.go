package users

import (
	"github.com/zengqiang96/mattermostclone/model"
	"github.com/zengqiang96/mattermostclone/services/cache"
	"github.com/zengqiang96/mattermostclone/store"
)

type UserService struct {
	store        store.UserStore
	sessionStore store.SessionStore
	sessionCache cache.Cache
	config       func() *model.Config
}
