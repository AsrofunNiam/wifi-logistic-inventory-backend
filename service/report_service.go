package service

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/gin-gonic/gin"
)

type ReportService interface {
	GenerateStockSummary(auth *auth.AccessDetails, startDate, endDate string, c *gin.Context) web.StockSummaryReport
	GenerateStockIn(auth *auth.AccessDetails, startDate, endDate string, c *gin.Context) []web.StockInResponse
	GenerateStockOut(auth *auth.AccessDetails, startDate, endDate string, c *gin.Context) []web.StockOutResponse
	GenerateLowStock(auth *auth.AccessDetails, c *gin.Context) []web.LowStockProduct
}
