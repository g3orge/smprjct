package storage

import (
	"errors"
	"sync"
)

var (
	ErrNoSuchKey = errors.New("no such key")
)

type MyStorage struct {
	sync.RWMutex
	m map[string]string
}

func New() *MyStorage {
	cache := MyStorage{
		m: make(map[string]string),
	}

	return &cache
}

func (store *MyStorage) Get(key string) (string, error) {
	store.RLock()
	defer store.RUnlock()

	val, ok := store.m[key]
	if !ok {
		return "", ErrNoSuchKey
	}

	return val, nil
}

func (store *MyStorage) Set(key, value string) {
	store.Lock()
	defer store.Unlock()

	store.m[key] = value
}

func (store *MyStorage) Delete(key string) error {
	store.Lock()
	defer store.Unlock()

	if _, ok := store.m[key]; !ok {
		return ErrNoSuchKey
	}

	delete(store.m, key)

	return nil
}
