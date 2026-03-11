package service

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/gin-gonic/gin"
)

type SupplierService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.SupplierResponse
	FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.SupplierResponse
	Create(auth *auth.AccessDetails, request *web.SupplierCreateRequest, c *gin.Context) web.SupplierResponse
	Update(auth *auth.AccessDetails, id uint, request *web.SupplierUpdateRequest, c *gin.Context) web.SupplierResponse
	Delete(auth *auth.AccessDetails, id uint, c *gin.Context)
}
