package controller

import (
	"net/http"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
)

type DashboardControllerImpl struct {
	DashboardService service.DashboardService
}

func NewDashboardController(dashboardService service.DashboardService) DashboardController {
	return &DashboardControllerImpl{
		DashboardService: dashboardService,
	}
}

func (controller *DashboardControllerImpl) GetStats(c *gin.Context, auth *auth.AccessDetails) {
	dashboardStats := controller.DashboardService.GetStats(auth, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Dashboard stats retrieved successfully",
		Data:    dashboardStats,
	}

	c.JSON(http.StatusOK, webResponse)
}
