package service

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/gin-gonic/gin"
)

type StockOutService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.StockOutResponse
	FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.StockOutResponse
	Create(auth *auth.AccessDetails, request *web.StockOutCreateRequest, c *gin.Context) web.StockOutResponse
	Update(auth *auth.AccessDetails, id uint, request *web.StockOutUpdateRequest, c *gin.Context) web.StockOutResponse
	Delete(auth *auth.AccessDetails, id uint, c *gin.Context)
}
