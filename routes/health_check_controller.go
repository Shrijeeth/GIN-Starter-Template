package routes

import (
	"mymodule/config"
	health_check "mymodule/services"

	"github.com/gin-gonic/gin"
)

func RegisterHealthCheckService(router *gin.Engine, configs config.Config) {
	routerGroup := router.Group("health-check")
	healthCheckService := health_check.GetHealthCheckServiceInstance(configs)

	routerGroup.GET("/", healthCheckService.HealthCheck)
}