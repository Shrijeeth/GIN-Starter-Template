package health_check

import (
	"log"
	"mymodule/config"
	"mymodule/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type healthCheckService struct {
	db *gorm.DB
}
type HealthCheckServiceInterface interface {
	HealthCheck(c *gin.Context)
}
func GetHealthCheckServiceInstance (configs config.Config) HealthCheckServiceInterface {
	return &healthCheckService{
		db: configs.Db,
	}
}

func (healthCheck *healthCheckService) HealthCheck (c *gin.Context) {
	healthCheckStatus := map[string]interface{}{
		"server": "Healthy",
	}

	healthCheckData := models.HealthCheckModel{
		Status: "Healthy",
	}

	_ = healthCheck.db.Transaction(func(tx *gorm.DB) error {
		insertedData := healthCheck.db.Create(&healthCheckData)
		if insertedData.Error != nil || insertedData.RowsAffected != 1 {
			log.Println("Application Error: ", insertedData.Error, insertedData.RowsAffected)
			healthCheckStatus["database"] = "Unhealthy"
		} else {
			healthCheckStatus["database"] = "Healthy"
		}
		healthCheck.db.Unscoped().Delete(&healthCheckData)
		return nil
	})

	c.JSON(http.StatusOK, healthCheckStatus)
}