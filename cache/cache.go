package cache

import (
	"time"
)

var (
	DefaultCacher = NewMemoryCacher()
)

type Cacher interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string, value interface{}) error
	Close()
}

func Set(key string, value interface{}, expiration time.Duration) error {
	return DefaultCacher.Set(key, value, expiration)
}

func Get(key string, value interface{}) error {
	return DefaultCacher.Get(key, value)
}

func Close() {
	DefaultCacher.Close()
}
