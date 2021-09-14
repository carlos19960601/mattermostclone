package users

import (
	"time"

	"github.com/zengqiang96/mattermostclone/model"
)

func (us *UserService) CreateSession(session *model.Session) (*model.Session, error) {
	session.Token = ""

	session, err := us.sessionStore.Save(session)
	if err != nil {
		return nil, err
	}

	us.AddSessionToCache(session)

	return session, nil
}

func (us *UserService) AddSessionToCache(session *model.Session) {
	us.sessionCache.SetWithExpiry(session.Token, session, time.Duration(int64(*us.config().ServiceSettings.SessionCacheInMinutes))*time.Minute)
}
