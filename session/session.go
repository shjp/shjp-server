package session

import (
	"time"
)

// Session interface
type Session interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	Delete(key string)
	Name() string
}

// Session store struct
type Store struct {
	name     string
	accessed time.Time
	value    map[string]interface{}
}

func NewStore(name string) *Store {
	return &Store{
		name:  name,
		value: make(map[string]interface{}, 0),
	}
}

func (s *Store) Set(key string, value interface{}) {
	s.accessed = time.Now()
	s.value[key] = value
}

func (s *Store) Get(key string) (interface{}, bool) {
	s.accessed = time.Now()
	val, ok := s.value[key]
	return val, ok
}

func (s *Store) Delete(key string) {
	s.accessed = time.Now()
	delete(s.value, key)
}

func (s *Store) Name() string {
	return s.name
}
