package service

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/gin-gonic/gin"
)

type ProductService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.ProductResponse
	FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.ProductResponse
	FindImage(auth *auth.AccessDetails, image string, c *gin.Context) []byte
	Create(auth *auth.AccessDetails, request *web.ProductCreateRequest, c *gin.Context) web.ProductResponse
	Update(auth *auth.AccessDetails, id uint, request *web.ProductUpdateRequest, c *gin.Context) web.ProductResponse
	Delete(auth *auth.AccessDetails, id uint, c *gin.Context)

	// Group transaction
	FindAllTransaction(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.TransactionResponse
	CreateTransaction(auth *auth.AccessDetails, request *web.TransactionCreateRequest, c *gin.Context) web.TransactionResponse
}
