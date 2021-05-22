package retrylayer

import "github.com/zengqiang96/mattermostclone/store"

type RetryLayer struct {
	store.Store
}

func New(childStore store.Store) *RetryLayer {
	newStore := RetryLayer{
		Store: childStore,
	}

	return &newStore
}
