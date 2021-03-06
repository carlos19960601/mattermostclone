package app

import (
	"errors"
	"net/http"
	"sync"

	"github.com/zengqiang96/mattermostclone/model"
	"github.com/zengqiang96/mattermostclone/store"
)

var userSessionPool = sync.Pool{
	New: func() interface{} {
		return &model.Session{}
	},
}

func (a *App) CreateSession(session *model.Session) (*model.Session, *model.AppError) {
	session, err := a.srv.userService.CreateSession(session)

	if err != nil {
		var invErr *store.ErrInvalidInput
		switch {
		case errors.As(err, &invErr):
			return nil, model.NewAppError("CreateSession", "app.session.save.existing.app_error", nil, invErr.Error(), http.StatusBadRequest)
		default:
			return nil, model.NewAppError("CreateSession", "app.session.save.app_error", nil, err.Error(), http.StatusInternalServerError)
		}
	}

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

func (a *App) createSessionForUserAccessToken(tokenString string) (*model.Session, *model.AppError) {
	return nil, nil
}
