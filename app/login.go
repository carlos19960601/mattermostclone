package app

import (
	"net/http"
	"time"

	"github.com/zengqiang96/mattermostclone/app/request"
	"github.com/zengqiang96/mattermostclone/model"
	"github.com/zengqiang96/mattermostclone/utils"
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

func (a *App) AttachSessionCookies(c *request.Context, w http.ResponseWriter, r *http.Request) {
	secure := false

	maxAge := *a.Config().ServiceSettings.SessionLengthWebInDays * 60 * 60 * 24
	domain := a.GetCookieDomain()
	subpath, _ := utils.GetSubpathFromConfig(a.Config())

	expiresAt := time.Unix(model.GetMillis()/1000, 0)
	sessionCookie := &http.Cookie{
		Name:     model.SESSION_COOKIE_TOKEN,
		Value:    c.Session().Token,
		Path:     subpath,
		MaxAge:   maxAge,
		Expires:  expiresAt,
		HttpOnly: true,
		Domain:   domain,
		Secure:   secure,
	}

	http.SetCookie(w, sessionCookie)
}
