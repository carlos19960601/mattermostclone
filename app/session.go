package app

import (
	"net/http"
	"sync"
	"time"

	"github.com/zengqiang96/mattermostclone/model"
)

var userSessionPool = sync.Pool{
	New: func() interface{} {
		return &model.Session{}
	},
}

func (a *App) CreateSession(session *model.Session) (*model.Session, *model.AppError) {
	session.Token = ""

	session, err := a.Srv().Store.Session().Save(session)
	if err != nil {

	}

	a.AddSessionToCache(session)
	return session, nil
}

func ReturnSessionToPool(session *model.Session) {
	if session != nil {
		userSessionPool.Put(session)
	}
}

func (a *App) GetSession(token string) (*model.Session, *model.AppError) {
	var session = userSessionPool.Get().(*model.Session)
	var err *model.AppError
	if err := a.Srv().sessionCache.Get(token, session); err == nil {

	}

	if session == nil || session.Id == "" {
		session, err = a.createSessionForUserAccessToken(token)
		if err != nil {
			return nil, model.NewAppError("GetSession", "api.context.invalid_token.error", map[string]interface{}{"Token": token, "Error": err.Error()}, "", http.StatusUnauthorized)
		}
	}

	return session, nil
}

func (a *App) AddSessionToCache(session *model.Session) {
	a.Srv().sessionCache.SetWithExpire(session.Token, session, time.Minute)
}

func (a *App) createSessionForUserAccessToken(tokenString string) (*model.Session, *model.AppError) {
	return nil, nil
}
