package main

import (
	"log"
	"os"

	routers "ShopOps/Delivery/routers"
	Infrastructure "ShopOps/Infrastructure"
)

func main() {
	// Initialize MongoDB
	if err := Infrastructure.InitMongo(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer Infrastructure.CloseMongo()

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Setup router using GetDB()
	router := routers.SetupRouter(Infrastructure.GetDB())

	// Start server
	log.Printf("ShopOps Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
