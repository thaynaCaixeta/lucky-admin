package main

import (
	"log"
)

func main() {
	app, err := injectAppDependencies()
	if err != nil {
		log.Fatalf("Failed to init app: %v", err)
	}
	// Start the application
	app.Run()
}

type Application interface {
	Run()
}
