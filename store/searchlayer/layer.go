package searchlayer

import "github.com/zengqiang96/mattermostclone/store"

type SearchStore struct {
	store.Store
}

func NewSearchLayer(baseStore store.Store) *SearchStore {
	searchStore := &SearchStore{
		Store: baseStore,
	}

	return searchStore
}
