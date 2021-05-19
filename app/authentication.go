package app

import (
	"net/http"

	"github.com/zengqiang96/mattermostclone/app/request"
	"github.com/zengqiang96/mattermostclone/model"
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
