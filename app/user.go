package app

import (
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
	return nil, nil
}
