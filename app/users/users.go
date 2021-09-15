package users

import (
	"github.com/pkg/errors"
	"github.com/zengqiang96/mattermostclone/model"
)

type UserCreateOptions struct {
	Guest bool
}

func (us *UserService) CreateUser(user *model.User, opts UserCreateOptions) (*model.User, error) {
	user.Roles = model.SystemUserRoleId
	if opts.Guest {
		user.Roles = model.SystemGuestRoleId
	}

	if ok, err := us.store.IsEmpty(true); err != nil {
		return nil, errors.Wrap(UserStoreIsEmptyError, err.Error())
	} else if ok {
		user.Roles = model.SystemAdminRoleId + " " + model.SystemUserRoleId
	}

	return us.createUser(user)
}

func (us *UserService) createUser(user *model.User) (*model.User, error) {
	ruser, err := us.store.Save(user)
	if err != nil {
		return nil, err
	}

	return ruser, nil
}
