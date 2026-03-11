package route

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/controller"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReportRoute(router *gin.Engine, db *gorm.DB) {
	reportService := service.NewReportService(db)
	reportController := controller.NewReportController(reportService)

	router.GET("/reports/stock-summary", auth.Auth(reportController.GenerateStockSummary, []string{}))
	router.GET("/reports/stock-in", auth.Auth(reportController.GenerateStockIn, []string{}))
	router.GET("/reports/stock-out", auth.Auth(reportController.GenerateStockOut, []string{}))
	router.GET("/reports/low-stock", auth.Auth(reportController.GenerateLowStock, []string{}))
}
