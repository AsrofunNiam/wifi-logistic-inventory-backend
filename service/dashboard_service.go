package service

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/gin-gonic/gin"
)

type DashboardService interface {
	GetStats(auth *auth.AccessDetails, c *gin.Context) web.DashboardStatsResponse
}
