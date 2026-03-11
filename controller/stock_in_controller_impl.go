package controller

import (
	"net/http"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
)

type StockInControllerImpl struct {
	StockInService service.StockInService
}

func NewStockInController(stockInService service.StockInService) StockInController {
	return &StockInControllerImpl{
		StockInService: stockInService,
	}
}

func (controller *StockInControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "code.like", "product_id.eq", "supplier_id.eq", "date.gte", "date.lte")
	stockInResponses := controller.StockInService.FindAll(auth, &filters, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(stockInResponses),
		Data:    stockInResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *StockInControllerImpl) FindByID(c *gin.Context, auth *auth.AccessDetails) {
	stockInIDParam := c.Param("id")
	stockInID := helper.StringToUint(stockInIDParam)

	stockInResponse := controller.StockInService.FindByID(auth, &stockInID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(stockInResponse),
		Data:    stockInResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *StockInControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {
	var request web.StockInCreateRequest
	err := c.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	stockInResponse := controller.StockInService.Create(auth, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Stock in created successfully",
		Data:    stockInResponse,
	}
	c.JSON(http.StatusCreated, webResponse)
}

func (controller *StockInControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	stockInIDParam := c.Param("id")
	stockInID := helper.StringToUint(stockInIDParam)

	var request web.StockInUpdateRequest
	err := c.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	stockInResponse := controller.StockInService.Update(auth, stockInID, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Stock in updated successfully",
		Data:    stockInResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *StockInControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	stockInIDParam := c.Param("id")
	stockInID := helper.StringToUint(stockInIDParam)

	controller.StockInService.Delete(auth, stockInID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Stock in deleted successfully",
		Data:    nil,
	}
	c.JSON(http.StatusOK, webResponse)
}
