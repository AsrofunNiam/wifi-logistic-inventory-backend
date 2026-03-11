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
	filters := helper.FilterFromQueryString(c, "name.like", "id.eq")
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

func (controller *ProductControllerImpl) FindImage(c *gin.Context, auth *auth.AccessDetails) {
	fileName := c.Param("image_name")
	fileResponse := controller.ProductService.FindImage(auth, fileName, c)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Write(fileResponse)
}

func (controller *ProductControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {
	form, _ := c.MultipartForm()

	companyCodeStr := c.PostForm("company_code")

	companyCodeUint := helper.StringToUint(companyCodeStr)
	request := web.ProductCreateRequest{
		Name:        c.PostForm("name"),
		Type:        c.PostForm("type"),
		CompanyCode: companyCodeUint,
		Description: c.PostForm("description"),
		Available:   c.PostForm("available") == "true",
		ImageFile:   form.File["image_file"],
	}

	productResponse := controller.ProductService.Create(auth, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Product updated successfully",
		Data:    productResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	productID := c.Param("id")
	form, _ := c.MultipartForm()

	productIDUint := helper.StringToUint(productID)
	request := web.ProductUpdateRequest{
		Name:        c.PostForm("name"),
		Type:        c.PostForm("type"),
		Description: c.PostForm("description"),
		Available:   c.PostForm("available") == "true",
		ImageName:   c.PostForm("image_name"),
		ImageFile:   form.File["image_file"],
	}

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

// Group transaction

func (controller *ProductControllerImpl) FindAllProductTransaction(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "name.like", "id.eq")
	transactionResponses := controller.ProductService.FindAllTransaction(auth, &filters, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(transactionResponses),
		Data:    transactionResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductControllerImpl) CreateProductTransaction(c *gin.Context, auth *auth.AccessDetails) {
	request := &web.TransactionCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	transactionResponse := controller.ProductService.CreateTransaction(auth, request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Transaction created successfully",
		Data:    transactionResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}
