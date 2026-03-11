package controller

import (
	"net/http"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
)

type StockOutControllerImpl struct {
	StockOutService service.StockOutService
}

func NewStockOutController(stockOutService service.StockOutService) StockOutController {
	return &StockOutControllerImpl{
		StockOutService: stockOutService,
	}
}

func (controller *StockOutControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "code.like", "product_id.eq", "destination.like", "date.gte", "date.lte")
	stockOutResponses := controller.StockOutService.FindAll(auth, &filters, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(stockOutResponses),
		Data:    stockOutResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *StockOutControllerImpl) FindByID(c *gin.Context, auth *auth.AccessDetails) {
	stockOutIDParam := c.Param("id")
	stockOutID := helper.StringToUint(stockOutIDParam)

	stockOutResponse := controller.StockOutService.FindByID(auth, &stockOutID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(stockOutResponse),
		Data:    stockOutResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *StockOutControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {
	var request web.StockOutCreateRequest
	err := c.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	stockOutResponse := controller.StockOutService.Create(auth, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Stock out created successfully",
		Data:    stockOutResponse,
	}
	c.JSON(http.StatusCreated, webResponse)
}

func (controller *StockOutControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	stockOutIDParam := c.Param("id")
	stockOutID := helper.StringToUint(stockOutIDParam)

	var request web.StockOutUpdateRequest
	err := c.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	stockOutResponse := controller.StockOutService.Update(auth, stockOutID, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Stock out updated successfully",
		Data:    stockOutResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *StockOutControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	stockOutIDParam := c.Param("id")
	stockOutID := helper.StringToUint(stockOutIDParam)

	controller.StockOutService.Delete(auth, stockOutID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Stock out deleted successfully",
		Data:    nil,
	}
	c.JSON(http.StatusOK, webResponse)
}
