package config

import "sync"

type BackingStore interface {
	Close() error
}

type Store struct {
	backingStore BackingStore

	configLock sync.RWMutex
}

func NewStore(dsn string) (*Store, error) {
	backingStore, err := getBackingStore(dsn)
	if err != nil {
		return nil, err
	}
	store, err := NewStoreFromBacking(backingStore)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func NewStoreFromBacking(backingStore BackingStore) (*Store, error) {
	store := Store{
		backingStore: backingStore,
	}

	return &store, nil
}

func getBackingStore(dsn string) (BackingStore, error) {
	return NewFileStore(dsn)
}

func (s *Store) Close() error {
	s.configLock.Lock()
	defer s.configLock.Unlock()
	return s.backingStore.Close()
}
