package localcachelayer

import "github.com/zengqiang96/mattermostclone/store"

type LocalCacheLayer struct {
	store.Store
}

func NewLocalCacheLayer(baseStore store.Store) (localcachelayer LocalCacheLayer, err error) {
	localcachelayer = LocalCacheLayer{
		Store: baseStore,
	}

	return localcachelayer, nil
}
