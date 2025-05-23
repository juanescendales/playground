package app

import (
	"fmt"
	"log"

	"github.com/juanescendales/playground/system-design/go-simple-cache/internal/infrastructure/adapters"
	"github.com/juanescendales/playground/system-design/go-simple-cache/internal/infrastructure/adapters/external/cache"
	"github.com/juanescendales/playground/system-design/go-simple-cache/internal/infrastructure/adapters/external/db"
	"github.com/juanescendales/playground/system-design/go-simple-cache/internal/infrastructure/entrypoint/rest"
)

// AppStart initializes and starts the application
// Returns an error if initialization fails
func AppStart() error {
	config, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	db := db.NewMockDB()
	cache := cache.NewLeastRecentlyUsedStrategy(config.CacheCapacity)
	repository := adapters.NewRepository(cache, db)
	handler := rest.NewHandler(repository)

	// Start the web server
	log.Println("Starting web server...")
	rest.Start(handler)

	return nil
}
