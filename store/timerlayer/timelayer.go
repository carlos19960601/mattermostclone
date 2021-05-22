package timerlayer

import "github.com/zengqiang96/mattermostclone/store"

type TimerLayer struct {
	store.Store
	UserStore store.UserStore
}

func New(childStore store.Store) *TimerLayer {
	newStore := TimerLayer{
		Store: childStore,
	}

	newStore.UserStore = &TimerLayerUserStore{UserStore: childStore.User()}
	return &newStore
}

type TimerLayerUserStore struct {
	store.UserStore
	Root *TimerLayer
}
