package app

import (
	"net/http"

	"github.com/zengqiang96/mattermostclone/app/request"
	"github.com/zengqiang96/mattermostclone/model"
)

func (a *App) AuthenticateUserForLogin(c *request.Context, loginId, password string) (user *model.User, err *model.AppError) {
	if user, err = a.GetUserForLogin(loginId); err != nil {
		return nil, err
	}

	if user, err = a.authenticateUser(c, user, password); err != nil {
		return nil, err
	}
	return user, nil
}

func (a *App) GetUserForLogin(loginId string) (*model.User, *model.AppError) {
	if user, err := a.Srv().Store.User().GetForLogin(loginId); err == nil {
		return user, nil
	}

	return nil, model.NewAppError("GetUserForLogin", "store.sql_user.get_for_login.app_error", nil, "", http.StatusBadRequest)
}

func (a *App) DoLogin(c *request.Context, w http.ResponseWriter, r *http.Request, user *model.User) *model.AppError {
	session := &model.Session{
		UserId: user.Id,
	}

	var err *model.AppError
	if session, err = a.CreateSession(session); err != nil {
		err.StatusCode = http.StatusInternalServerError
		return err
	}

	w.Header().Set(model.HEADER_TOKEN, session.Token)

	c.SetSession(session)
	return nil
}
