package cache

import "time"

type Cache interface {
	SetWithExpiry(key string, value interface{}, ttl time.Duration) error
	Get(key string, value interface{}) error
}
