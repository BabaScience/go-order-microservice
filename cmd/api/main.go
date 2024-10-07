package main

import (
	"goorder/internal/api"
	"goorder/internal/db"
	"goorder/internal/models"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Initialize the database
	err = db.InitDB()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	// Migrate the models
	err = models.MigrateSocieta(db.DB)
	if err != nil {
		log.Fatalf("Failed to migrate Societa model: %v", err)
	}
	err = models.MigratePosizione(db.DB)
	if err != nil {
		log.Fatalf("Failed to migrate Deposito model: %v", err)
	}
	err = models.MigrateOrder(db.DB)
	if err != nil {
		log.Fatalf("Failed to migrate Order model: %v", err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	// Set up the router
	router := api.SetupRouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/build/")))

	handler := c.Handler(router)

	// Start the server
	log.Println("API server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
