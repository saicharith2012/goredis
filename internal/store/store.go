package store

import (
	"errors"
	"sync"
)

type ValueType string

var ErrWrongType = errors.New("wrong type")
var ErrInvalidInteger = errors.New("invalid integer")
var ErrNotFound = errors.New("not found")

const (
	String ValueType = "string"
	List   ValueType = "list"
)

type ValueWithExpiry struct {
	valueType ValueType
	data      any
	expiry    int64
}

type Store struct {
	mu     sync.Mutex
	values map[string]ValueWithExpiry
}

func New() *Store {
	return &Store{values: make(map[string]ValueWithExpiry)}
}
