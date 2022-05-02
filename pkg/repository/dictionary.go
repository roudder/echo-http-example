package repository

import "sync"

type Dictionary struct {
	mu sync.RWMutex
	m  map[string]string
}

type DictRepository interface {
	Get(key string) string
	Set(key string, value string)
}

func (d *Dictionary) Get(key string) string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.m[key]
}

func (d *Dictionary) Set(key string, value string){
	d.mu.Lock()
	d.m[key] = value
	d.mu.Unlock()
}

func NewDictionary() DictRepository {
	return &Dictionary{
		mu:   sync.RWMutex{},
		m: make(map[string]string),
	}
}
