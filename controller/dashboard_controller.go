package controller

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/gin-gonic/gin"
)

type DashboardController interface {
	GetStats(context *gin.Context, auth *auth.AccessDetails)
}
