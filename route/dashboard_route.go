package route

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/controller"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func DashboardRoute(router *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	dashboardService := service.NewDashboardService(
		db,
		redisClient,
	)
	dashboardController := controller.NewDashboardController(dashboardService)

	router.GET("/dashboard/stats", auth.Auth(dashboardController.GetStats, []string{}))
}
