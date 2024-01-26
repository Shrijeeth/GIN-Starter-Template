package routes

import (
	"mymodule/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes (router *gin.Engine, configs config.Config) {
	RegisterHealthCheckService(router, configs)
}