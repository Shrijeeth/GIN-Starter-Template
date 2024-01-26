package main

import (
	"log"
	"mymodule/config"
	"mymodule/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main () {
	log.Println("Starting server...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables")
	}
	log.Println("Environment variables loaded")

	configs := config.InitializeConfigs()
	log.Println("Configs loaded")

	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode("debug")
	router := gin.Default()
	routes.RegisterRoutes(router, configs)

	err = router.Run(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
