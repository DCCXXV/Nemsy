package auth

import (
	"sync"
	"time"
)

type StateStore struct {
	m   map[string]time.Time
	mu  sync.Mutex
	ttl time.Duration
}

func NewStateStore(ttl time.Duration) *StateStore {
	return &StateStore{
		m:   make(map[string]time.Time),
		ttl: ttl,
	}
}

func (s *StateStore) Put(state string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[state] = time.Now().Add(s.ttl)
}

func (s *StateStore) Check(state string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	exp, ok := s.m[state]
	if !ok {
		return false
	}
	delete(s.m, state)
	return time.Now().Before(exp)
}
