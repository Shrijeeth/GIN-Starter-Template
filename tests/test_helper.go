package tests

import (
	"log"
	"mymodule/config"
	"mymodule/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupTestRouter() (*gin.Engine, *config.Config) {
	log.Println("Starting test server...")

	err := godotenv.Load("../../.test.env")
	if err != nil {
		log.Fatalf("Error loading environment variables")
	}
	log.Println("Environment variables loaded")

	configs := config.InitializeConfigs()
	log.Println("Configs loaded")

	gin.SetMode("test")
	router := gin.Default()
	routes.RegisterRoutes(router, configs)
	return router, &configs
}