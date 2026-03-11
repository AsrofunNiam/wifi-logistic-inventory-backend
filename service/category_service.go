package service

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.CategoryResponse
	FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.CategoryResponse
	Create(auth *auth.AccessDetails, request *web.CategoryCreateRequest, c *gin.Context) web.CategoryResponse
	Update(auth *auth.AccessDetails, id uint, request *web.CategoryUpdateRequest, c *gin.Context) web.CategoryResponse
	Delete(auth *auth.AccessDetails, id uint, c *gin.Context)
}
