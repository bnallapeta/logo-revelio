package main

import (
	"log"
	"net/http"

	"github.com/bnallapeta/logo-revelio/api"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set up the Gin router
	r := gin.Default()

	// Set up routes
	api.SetupRoutes(r)

	// Start the server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
