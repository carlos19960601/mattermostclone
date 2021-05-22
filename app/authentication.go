package app

import (
	"net/http"
	"strings"

	"github.com/zengqiang96/mattermostclone/app/request"
	"github.com/zengqiang96/mattermostclone/model"
)

type TokenLocation int

const (
	TokenLocationNotFound TokenLocation = iota
	TokenLocationHeader
	TokenLocationCookie
	TokenLocationQueryString
	TokenLocationCloudHeader
	TokenLocationRemoteClusterHeader
)

func (a *App) authenticateUser(c *request.Context, user *model.User, password string) (*model.User, *model.AppError) {
	if err := a.CheckPasswordAndAllCriteria(user, password); err != nil {
		err.StatusCode = http.StatusUnauthorized
		return user, err
	}
	return user, nil
}

func (a *App) CheckPasswordAndAllCriteria(user *model.User, password string) *model.AppError {
	if err := a.checkUserPassword(user, password); err != nil {
		return err
	}

	return nil
}

func (a *App) checkUserPassword(user *model.User, password string) *model.AppError {
	if !model.ComparePassword(user.Password, password) {
		return model.NewAppError("checkUserPassword", "api.user.check_user_password.invalid.app_error", nil, "user_id"+user.Id, http.StatusUnauthorized)
	}
	return nil
}

func ParseAuthTokenFromRequest(r *http.Request) (string, TokenLocation) {
	authHeader := r.Header.Get(model.HEADER_AUTH)

	if cookie, err := r.Cookie(model.SESSION_COOKIE_TOKEN); err == nil {
		return cookie.Value, TokenLocationCookie
	}

	if len(authHeader) > 6 && strings.ToUpper(authHeader[0:6]) == model.HEADER_BEARER {
		return authHeader[7:], TokenLocationHeader
	}

	if len(authHeader) > 5 && strings.ToLower(authHeader[0:5]) == model.HEADER_TOKEN {
		return authHeader[6:], TokenLocationHeader
	}

	return "", TokenLocationNotFound
}
