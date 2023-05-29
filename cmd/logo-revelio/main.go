package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bnallapeta/logo-revelio/api"
	"github.com/bnallapeta/logo-revelio/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create the data directory if it doesn't exist
	if err := os.MkdirAll("../../data", 0755); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	// Initialize the database
	db, err := store.InitializeDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	// Initialize the LogosMap
	if err := store.LoadLogosMap(); err != nil {
		log.Fatal("Failed to load logos map:", err)
	}
	// Set up the Gin router
	r := gin.Default()

	// Set up routes
	api.SetupRoutes(r, db)

	// Start the server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
