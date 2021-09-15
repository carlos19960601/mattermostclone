package users

import "github.com/zengqiang96/mattermostclone/model"

func (us *UserService) IsFirstUserAccount() bool {
	count, err := us.store.Count(model.UserCountOptions{IncludeDeleted: true})
	if err != nil {
		return false
	}
	if count <= 0 {
		return true
	}
	return false
}
