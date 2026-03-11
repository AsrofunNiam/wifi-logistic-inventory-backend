package controller

import (
	"net/http"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "name.like")
	categoryResponses := controller.CategoryService.FindAll(auth, &filters, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(categoryResponses),
		Data:    categoryResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryControllerImpl) FindByID(c *gin.Context, auth *auth.AccessDetails) {
	categoryIDParam := c.Param("id")
	categoryID := helper.StringToUint(categoryIDParam)

	categoryResponse := controller.CategoryService.FindByID(auth, &categoryID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(categoryResponse),
		Data:    categoryResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {
	var request web.CategoryCreateRequest
	err := c.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Create(auth, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Category created successfully",
		Data:    categoryResponse,
	}
	c.JSON(http.StatusCreated, webResponse)
}

func (controller *CategoryControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	categoryIDParam := c.Param("id")
	categoryID := helper.StringToUint(categoryIDParam)

	var request web.CategoryUpdateRequest
	err := c.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Update(auth, categoryID, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Category updated successfully",
		Data:    categoryResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	categoryIDParam := c.Param("id")
	categoryID := helper.StringToUint(categoryIDParam)

	controller.CategoryService.Delete(auth, categoryID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Category deleted successfully",
		Data:    nil,
	}
	c.JSON(http.StatusOK, webResponse)
}
