package store

import (
	"strconv"
	"time"
)

func (s *Store) Set(key, value string, ttlSeconds int64) {
	var expiryTime int64
	if ttlSeconds > 0 {
		expiryTime = time.Now().Unix() + ttlSeconds
	} else {
		expiryTime = 0
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key] = ValueWithExpiry{data: value, expiry: expiryTime, valueType: String}
}

func (s *Store) Get(key string) (data string, isFound bool, err error) {

	s.mu.Lock()
	defer s.mu.Unlock()
	s.checkAndDeleteIfExpired(key)
	value, found := s.values[key]

	if !found {
		return "", false, nil
	}

	if value.valueType != String {
		return "", true, ErrWrongType
	}

	return value.data.(string), true, nil
}

func (s *Store) Incr(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.checkAndDeleteIfExpired(key)

	value, ok := s.values[key]

	if !ok {
		s.values[key] = ValueWithExpiry{valueType: String, data: "1", expiry: 0}
		return 1, nil
	}

	if value.valueType != String {
		return -1, ErrWrongType
	}

	val, err := strconv.Atoi(value.data.(string))

	if err != nil {
		return -1, ErrInvalidInteger
	}

	val += 1
	s.values[key] = ValueWithExpiry{data: strconv.Itoa(val), valueType: s.values[key].valueType, expiry: s.values[key].expiry}

	return val, nil

}
