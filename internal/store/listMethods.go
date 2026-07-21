package store

import (
	"fmt"
	"math"
)

// list commands
func (s *Store) LPush(key string, values []string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newList := []string{}
	s.checkAndDeleteIfExpired(key)

	for _, value := range values {
		newList = append([]string{value}, newList...)
	}

	value, ok := s.values[key]

	if !ok {
		s.values[key] = ValueWithExpiry{valueType: List, data: newList, expiry: 0}
		return len(values), nil
	}

	if value.valueType != List {
		return 0, ErrWrongType
	}

	list := value.data.([]string)
	list = append(newList, list...)

	s.values[key] = ValueWithExpiry{data: list, valueType: s.values[key].valueType, expiry: s.values[key].expiry}

	return len(list), nil

}

func (s *Store) RPush(key string, values []string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.checkAndDeleteIfExpired(key)

	value, ok := s.values[key]

	if !ok {
		s.values[key] = ValueWithExpiry{valueType: List, data: values, expiry: 0}
		return len(values), nil
	}

	if value.valueType != List {
		return 0, ErrWrongType
	}

	list := value.data.([]string)
	list = append(list, values...)

	s.values[key] = ValueWithExpiry{data: list, valueType: s.values[key].valueType, expiry: s.values[key].expiry}

	fmt.Println(list)

	return len(list), nil
}

func (s *Store) LLen(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.checkAndDeleteIfExpired(key)

	value, ok := s.values[key]

	if !ok {
		return 0, nil
	}

	if value.valueType != List {
		return 0, ErrWrongType
	}

	return len(value.data.([]string)), nil

}

func (s *Store) LRange(key string, start int, end int) ([]string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.checkAndDeleteIfExpired(key)

	value, ok := s.values[key]

	if !ok {
		return []string{}, nil
	}

	if value.valueType != List {
		return nil, ErrWrongType
	}

	listLen := len(value.data.([]string))

	if start < 0 {
		start = listLen + start
	}

	if end < 0 {
		end = listLen + end
	}

	if math.Abs(float64(start)) > float64(listLen) {
		if start < 0 {
			start = 0
		} else {
			start = listLen
		}
	}

	if math.Abs(float64(end)) > float64(listLen) {
		if end > 0 {
			end = listLen - 1
		} else {
			end = -1
		}
	}

	if start > end {
		return []string{}, nil
	}

	result := value.data.([]string)[start : end+1]

	return result, nil
}

func (s *Store) LPop(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.checkAndDeleteIfExpired(key)

	value, ok := s.values[key]

	if !ok {
		return "", ErrNotFound
	}

	if value.valueType != List {
		return "", ErrWrongType
	}

	poppedValue := value.data.([]string)[0]
	newList := value.data.([]string)[1:]

	s.values[key] = ValueWithExpiry{data: newList, valueType: s.values[key].valueType, expiry: s.values[key].expiry}

	if len(s.values[key].data.([]string)) == 0 {
		delete(s.values, key)
	}

	return poppedValue, nil
}

func (s *Store) RPop(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.checkAndDeleteIfExpired(key)

	value, ok := s.values[key]

	if !ok {
		return "", ErrNotFound
	}

	if value.valueType != List {
		return "", ErrWrongType
	}

	list := value.data.([]string)

	poppedValue := list[len(list) - 1]
	newList := value.data.([]string)[:len(list) - 1]

	s.values[key] = ValueWithExpiry{data: newList, valueType: s.values[key].valueType, expiry: s.values[key].expiry}

	if len(s.values[key].data.([]string)) == 0 {
		delete(s.values, key)
	}

	return poppedValue, nil
}
