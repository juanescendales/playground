package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(handler *Handler) {
	// Set up routes
	http.HandleFunc("/key", handler.Handle)
	http.HandleFunc("/status", handler.Handle)

	// Create a new server
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil, // Use default ServeMux
	}

	// Channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	// Start the server in a goroutine
	go func() {
		log.Println("Server started on port 8080")
		serverErrors <- server.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Block until we receive a signal or an error
	select {
	case err := <-serverErrors:
		log.Printf("Error starting server: %v", err)
		return

	case <-shutdown:
		log.Println("Starting graceful shutdown...")

		// Create a deadline for the shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Attempt to gracefully shut down the server
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Could not gracefully shutdown the server: %v\n", err)

			// If graceful shutdown fails, forcefully close
			if err := server.Close(); err != nil {
				log.Printf("Could not close server: %v\n", err)
			}
		}
	}

	log.Println("Server stopped")
}
