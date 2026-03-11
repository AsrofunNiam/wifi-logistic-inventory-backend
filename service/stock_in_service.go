package service

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/gin-gonic/gin"
)

type StockInService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.StockInResponse
	FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.StockInResponse
	Create(auth *auth.AccessDetails, request *web.StockInCreateRequest, c *gin.Context) web.StockInResponse
	Update(auth *auth.AccessDetails, id uint, request *web.StockInUpdateRequest, c *gin.Context) web.StockInResponse
	Delete(auth *auth.AccessDetails, id uint, c *gin.Context)
}
