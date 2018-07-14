package session

import (
	"sync"
	"time"
)

// Session Manager struct
type Manager struct {
	lock        sync.Mutex
	maxLifeTime int64
}

var sessions = make(map[string]*Store, 0)

func NewManager(maxLifeTime int64) *Manager {
	return &Manager{maxLifeTime: maxLifeTime}
}

func (m *Manager) InitSession(sessionName string) *Store {
	session := NewStore(sessionName)
	sessions[sessionName] = session
	return session
}

func (m *Manager) Set(sessionName, key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	session, ok := sessions[sessionName]
	if !ok {
		session = m.InitSession(sessionName)
	}
	session.Set(key, value)
}

func (m *Manager) Get(sessionName, key string) (interface{}, bool) {
	session, ok := sessions[sessionName]
	if !ok {
		return nil, false
	}
	return session.Get(key)
}

func (m *Manager) SetExpire() {
	time.AfterFunc(time.Duration(m.maxLifeTime), func() {
		m.SetExpire()
	})
}
