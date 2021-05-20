package app

import (
	"net/http"

	"github.com/zengqiang96/mattermostclone/app/request"
	"github.com/zengqiang96/mattermostclone/model"
)

func (a *App) CreateUserFromSignup(c *request.Context, user *model.User) (*model.User, *model.AppError) {
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
	user.Roles = model.SYSTEM_USER_ROLE_ID
	if guest {
		user.Roles = model.SYSTEM_GUEST_ROLE_ID
	}

	count, err := a.Srv().Store.User().Count(model.UserCountOptions{IncludeDeleted: true})
	if err != nil {
		return nil, model.NewAppError("createUserOrGuest", "app.user.get_total_users_count.app_error", nil, err.Error(), http.StatusInternalServerError)
	}
	if count <= 0 {
		user.Roles = model.SYSTEM_ADMIN_ROLE_ID + " " + model.SYSTEM_USER_ROLE_ID
	}

	ruser, appErr := a.createUser(user)
	if appErr != nil {
		return nil, appErr
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
