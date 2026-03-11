package route

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/controller"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DashboardRoute(router *gin.Engine, db *gorm.DB) {
	dashboardService := service.NewDashboardService(
		db,
	)
	dashboardController := controller.NewDashboardController(dashboardService)

	router.GET("/dashboard/stats", auth.Auth(dashboardController.GetStats, []string{}))
}
