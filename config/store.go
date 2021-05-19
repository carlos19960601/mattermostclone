package config

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/zengqiang96/mattermostclone/model"
)

type BackingStore interface {
	Load() ([]byte, error)
	Close() error
}

type Store struct {
	backingStore BackingStore

	configLock sync.RWMutex
	config     *model.Config
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

	if err := store.Load(); err != nil {
		return nil, fmt.Errorf("加载配置失败 err: %w", err)
	}

	return &store, nil
}

func getBackingStore(dsn string) (BackingStore, error) {
	return NewFileStore(dsn)
}

func (s *Store) Load() error {
	s.configLock.Lock()
	var unlockOnce sync.Once
	defer unlockOnce.Do(s.configLock.Unlock)

	return s.loadLockedWithOld()
}

func (s *Store) Get() *model.Config {
	s.configLock.RLock()
	defer s.configLock.RUnlock()
	return s.config
}

func (s *Store) Close() error {
	s.configLock.Lock()
	defer s.configLock.Unlock()
	return s.backingStore.Close()
}

func (s *Store) loadLockedWithOld() error {
	configBytes, err := s.backingStore.Load()
	if err != nil {
		return err
	}

	loadedConfig := model.Config{}
	if len(configBytes) != 0 {
		if err = json.Unmarshal(configBytes, &loadedConfig); err != nil {
			return fmt.Errorf("反序列化配置内容失败 err: %w", err)
		}
	}

	s.config = &loadedConfig
	return nil
}
