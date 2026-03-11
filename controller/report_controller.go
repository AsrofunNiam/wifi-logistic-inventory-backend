package controller

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/gin-gonic/gin"
)

type ReportController interface {
	GenerateStockSummary(context *gin.Context, auth *auth.AccessDetails)
	GenerateStockIn(context *gin.Context, auth *auth.AccessDetails)
	GenerateStockOut(context *gin.Context, auth *auth.AccessDetails)
	GenerateLowStock(context *gin.Context, auth *auth.AccessDetails)
}
