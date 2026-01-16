package store

import "sync"

type SharedState struct {
	mu   sync.Mutex
	data map[string]string
}

func NewSharedState() *SharedState {
	return &SharedState{
		data: make(map[string]string),
	}
}

func (s *SharedState) SetValue(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *SharedState) GetValue(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.data[key]
	return value, ok
}
