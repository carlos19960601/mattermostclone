package app

import (
	"net/http"

	"github.com/zengqiang96/mattermostclone/app/request"
	"github.com/zengqiang96/mattermostclone/app/users"
	"github.com/zengqiang96/mattermostclone/model"
)

func (a *App) CreateUserFromSignup(c *request.Context, user *model.User) (*model.User, *model.AppError) {
	if !a.IsFirstUserAccount() {
		err := model.NewAppError("CreateUserFromSignup", "api.user.create_user.no_open_server", nil, "email="+user.Email, http.StatusForbidden)
		return nil, err
	}

	ruser, err := a.CreateUser(c, user)
	if err != nil {
		return nil, err
	}
	return ruser, nil
}

func (a *App) CreateUser(c *request.Context, user *model.User) (*model.User, *model.AppError) {
	return a.createUserOrGuest(c, user, false)
}

func (a *App) createUserOrGuest(c *request.Context, user *model.User, guest bool) (*model.User, *model.AppError) {
	ruser, nErr := a.srv.userService.CreateUser(user, users.UserCreateOptions{Guest: guest})
	if nErr != nil {
		return nil, model.NewAppError("createUserOrGuest", "app.user.save.app_error", nil, nErr.Error(), http.StatusInternalServerError)
	}
	return ruser, nil
}

func (a *App) createUser(user *model.User) (*model.User, *model.AppError) {
	ruser, err := a.Srv().Store.User().Save(user)
	if err != nil {
		return nil, model.NewAppError("createUser", "app.user.save.app_error", nil, err.Error(), http.StatusInternalServerError)
	}
	return ruser, nil
}

func (a *App) IsFirstUserAccount() bool {
	return a.srv.userService.IsFirstUserAccount()
}
