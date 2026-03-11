package controller

import (
	"net/http"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
)

type SupplierControllerImpl struct {
	SupplierService service.SupplierService
}

func NewSupplierController(supplierService service.SupplierService) SupplierController {
	return &SupplierControllerImpl{
		SupplierService: supplierService,
	}
}

func (controller *SupplierControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "name.like", "code.like", "status.eq")
	supplierResponses := controller.SupplierService.FindAll(auth, &filters, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(supplierResponses),
		Data:    supplierResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *SupplierControllerImpl) FindByID(c *gin.Context, auth *auth.AccessDetails) {
	supplierIDParam := c.Param("id")
	supplierID := helper.StringToUint(supplierIDParam)

	supplierResponse := controller.SupplierService.FindByID(auth, &supplierID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(supplierResponse),
		Data:    supplierResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *SupplierControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {
	var request web.SupplierCreateRequest
	err := c.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	supplierResponse := controller.SupplierService.Create(auth, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Supplier created successfully",
		Data:    supplierResponse,
	}
	c.JSON(http.StatusCreated, webResponse)
}

func (controller *SupplierControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	supplierIDParam := c.Param("id")
	supplierID := helper.StringToUint(supplierIDParam)

	var request web.SupplierUpdateRequest
	err := c.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	supplierResponse := controller.SupplierService.Update(auth, supplierID, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Supplier updated successfully",
		Data:    supplierResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *SupplierControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	supplierIDParam := c.Param("id")
	supplierID := helper.StringToUint(supplierIDParam)

	controller.SupplierService.Delete(auth, supplierID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Supplier deleted successfully",
		Data:    nil,
	}
	c.JSON(http.StatusOK, webResponse)
}
