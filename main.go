package main

import (
	"fproj/models"
	"fproj/utils"
	"fproj/controllers"
	"fproj/routes"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
    // "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	utils.Setenv()
	// Create a new gin instance
	r := gin.Default()

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
    if err != nil {
        log.Fatal(err)
    }
    config := models.Config{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
    }

    // Initialize DB
    db := models.InitDB(config)
	
	controllers.SetDB(db)
	// Load the routes
	routes.AuthRoutes(r)

	// Run the server
	r.Run(":8080")
	print("hehehihi")
}