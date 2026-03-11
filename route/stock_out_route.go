package route

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/controller"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/repository"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func StockOutRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {
	stockOutService := service.NewStockOutService(
		repository.NewStockOutRepository(),
		repository.NewProductRepository(),
		db,
		validate,
	)
	stockOutController := controller.NewStockOutController(stockOutService)

	router.GET("/stock-out", auth.Auth(stockOutController.FindAll, []string{}))
	router.GET("/stock-out/:id", auth.Auth(stockOutController.FindByID, []string{}))
	router.POST("/stock-out", auth.Auth(stockOutController.Create, []string{}))
	router.PUT("/stock-out/:id", auth.Auth(stockOutController.Update, []string{}))
	router.DELETE("/stock-out/:id", auth.Auth(stockOutController.Delete, []string{}))
}
