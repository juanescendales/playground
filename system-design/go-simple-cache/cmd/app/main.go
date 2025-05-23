package main

import (
	"fmt"
	"log"
	"os"

	"github.com/juanescendales/playground/system-design/go-simple-cache/internal/app"
)

func main() {
	fmt.Println("Starting Go Simple Cache")

	if err := app.AppStart(); err != nil {
		log.Printf("Application failed to start: %v", err)
		os.Exit(1)
	}
}
