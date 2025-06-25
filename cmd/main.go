package main

import (
	"hris-backend/config"
	"hris-backend/routes"
	"hris-backend/services"
	"log"
)

func main() {
	// Initialize database
	config.InitDB()

	// Run database migrations
	services.RunMigrations(config.DB)

	// Create Fiber app
	app := config.NewFiberApp()

	// Setup routes
	routes.SetupRoutes(app)

	// Get port from config
	port := config.GetPort()

	// Start server
	log.Printf("Server starting on %s", port)
	log.Fatal(app.Listen(port))
}
