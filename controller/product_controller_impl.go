package controller

import (
	"net/http"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "name.like", "id.eq", "category_id.eq", "supplier_id.eq")
	productResponses := controller.ProductService.FindAll(auth, &filters, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(productResponses),
		Data:    productResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductControllerImpl) FindByID(c *gin.Context, auth *auth.AccessDetails) {
	productIDParam := c.Param("id")
	productID := helper.StringToUint(productIDParam)

	productResponse := controller.ProductService.FindByID(auth, &productID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(productResponse),
		Data:    productResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {
	var request web.ProductCreateRequest
	helper.ReadFromRequestBody(c, &request)

	productResponse := controller.ProductService.Create(auth, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Product created successfully",
		Data:    productResponse,
	}
	c.JSON(http.StatusCreated, webResponse)
}

func (controller *ProductControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	productID := c.Param("id")
	productIDUint := helper.StringToUint(productID)

	var request web.ProductUpdateRequest
	helper.ReadFromRequestBody(c, &request)

	productResponse := controller.ProductService.Update(auth, productIDUint, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Product updated successfully",
		Data:    productResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	productID := c.Param("id")
	productIDUint := helper.StringToUint(productID)

	controller.ProductService.Delete(auth, productIDUint, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Product deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}
