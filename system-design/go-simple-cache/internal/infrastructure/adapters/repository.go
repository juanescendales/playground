package adapters

import "log"

type Cache interface {
	Get(key string) ([]byte, error)
	Add(key string, value []byte) error
	OrederedKeys() []string
	Size() int
}
type Database interface {
	Get(key string) ([]byte, error)
	Add(key string, value []byte) error
}

type Repository struct {
	cache    Cache
	database Database
}

func NewRepository(cache Cache, database Database) *Repository {
	return &Repository{
		cache:    cache,
		database: database,
	}
}

func (r *Repository) Get(key string) ([]byte, error) {
	value, err := r.cache.Get(key)
	if err == nil {
		return value, nil
	}
	log.Printf("Cache miss for key: %s", key)
	value, err = r.database.Get(key)
	if err != nil {
		return nil, err
	}
	r.cache.Add(key, value)
	return value, nil
}

func (r *Repository) Add(key string, value []byte) error {
	return r.database.Add(key, value)
}

func (r *Repository) CacheStatus() *CacheStatus {
	return &CacheStatus{
		OrderedKeys: r.cache.OrederedKeys(),
		Size:        r.cache.Size(),
	}
}

type CacheStatus struct {
	OrderedKeys []string
	Size        int
}
