package users

import "errors"

var (
	UserStoreIsEmptyError = errors.New("could not check if the user store is empty")
)
