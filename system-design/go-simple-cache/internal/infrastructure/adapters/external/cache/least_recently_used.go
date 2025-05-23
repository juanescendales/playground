package cache

import (
	"errors"
	"log"

	"github.com/juanescendales/playground/system-design/go-simple-cache/internal/infrastructure/adapters/external/collection"
)

var ErrKeyNotFound = errors.New("key not found")

type LeastRecentlyUsedStrategy struct {
	cacheMap map[string]*collection.Node
	list     *collection.LinkedList
	capacity int
}

func NewLeastRecentlyUsedStrategy(capacity int) *LeastRecentlyUsedStrategy {
	return &LeastRecentlyUsedStrategy{
		cacheMap: make(map[string]*collection.Node),
		list:     collection.New(),
		capacity: capacity,
	}
}

func (s *LeastRecentlyUsedStrategy) Get(key string) ([]byte, error) {
	if _, ok := s.cacheMap[key]; ok {
		s.list.ToHead(s.cacheMap[key])
		return s.cacheMap[key].Value(), nil
	}
	return nil, ErrKeyNotFound
}

func (s *LeastRecentlyUsedStrategy) Add(key string, value []byte) error {
	if _, ok := s.cacheMap[key]; ok {
		s.list.ToHead(s.cacheMap[key])
		return nil
	}

	if s.list.Size() == s.capacity {
		node := s.list.Pop()
		s.cacheMap[node.Key()] = nil
		log.Printf("Removed key: %s from cache", node.Key())
	}

	node := collection.NewNode(key, value)
	s.cacheMap[key] = node
	s.list.Push(node)

	return nil
}

func (s *LeastRecentlyUsedStrategy) OrederedKeys() []string {
	return s.list.OrderedKeys()
}

func (s *LeastRecentlyUsedStrategy) Size() int {
	return s.list.Size()
}
