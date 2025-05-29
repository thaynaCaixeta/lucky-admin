package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	app, err := injectAppDependencies(ctx)
	if err != nil {
		log.Fatalf("Failed to init app: %v", err)
	}
	// Start the application
	app.Run()
}

type Application interface {
	Run()
}
