package store

import (
	"time"
)

func (s *Store) Exists(keys []string) int {
	exists := 0
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, key := range keys {
		s.checkAndDeleteIfExpired(key)
		if _, ok := s.values[key]; ok {
			exists++
		}
	}

	return exists
}

func (s *Store) Type(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.checkAndDeleteIfExpired(key)
	value, ok := s.values[key]

	if !ok {
		return "none"
	}

	return string(value.valueType)
}

func (s *Store) Delete(keys []string) int {
	deletes := 0

	s.mu.Lock()
	defer s.mu.Unlock()
	for _, key := range keys {
		s.checkAndDeleteIfExpired(key)
		if _, ok := s.values[key]; ok {
			delete(s.values, key)
			deletes++
		}
	}

	return deletes
}

func (s *Store) TTL(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.checkAndDeleteIfExpired(key)
	value, ok := s.values[key]

	if !ok {
		return -2
	}

	if value.expiry == 0 {
		return -1
	}

	return int(value.expiry - time.Now().Unix())
}

// expiry checker
func (s *Store) checkAndDeleteIfExpired(key string) {
	if value, ok := s.values[key]; ok && value.expiry != 0 && value.expiry <= time.Now().Unix() {
		delete(s.values, key)
	}
}
