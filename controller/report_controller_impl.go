package controller

import (
	"net/http"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
)

type ReportControllerImpl struct {
	ReportService service.ReportService
}

func NewReportController(reportService service.ReportService) ReportController {
	return &ReportControllerImpl{
		ReportService: reportService,
	}
}

func (controller *ReportControllerImpl) GenerateStockSummary(c *gin.Context, auth *auth.AccessDetails) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	report := controller.ReportService.GenerateStockSummary(auth, startDate, endDate, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Stock summary report generated successfully",
		Data:    report,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ReportControllerImpl) GenerateStockIn(c *gin.Context, auth *auth.AccessDetails) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	report := controller.ReportService.GenerateStockIn(auth, startDate, endDate, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Stock in report generated successfully",
		Data:    report,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ReportControllerImpl) GenerateStockOut(c *gin.Context, auth *auth.AccessDetails) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	report := controller.ReportService.GenerateStockOut(auth, startDate, endDate, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Stock out report generated successfully",
		Data:    report,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ReportControllerImpl) GenerateLowStock(c *gin.Context, auth *auth.AccessDetails) {
	report := controller.ReportService.GenerateLowStock(auth, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Low stock report generated successfully",
		Data:    report,
	}

	c.JSON(http.StatusOK, webResponse)
}
