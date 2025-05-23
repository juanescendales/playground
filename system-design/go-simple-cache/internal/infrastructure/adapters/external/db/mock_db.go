package db

import (
	"errors"
)

var ErrKeyNotFound = errors.New("key not found")

type MockDB struct {
	store map[string][]byte
}

func NewMockDB() *MockDB {
	return &MockDB{
		store: make(map[string][]byte),
	}
}

func (m *MockDB) Get(key string) ([]byte, error) {
	if value, exists := m.store[key]; exists {
		return value, nil
	}
	return nil, ErrKeyNotFound
}

func (m *MockDB) Add(key string, value []byte) error {
	m.store[key] = value
	return nil
}
