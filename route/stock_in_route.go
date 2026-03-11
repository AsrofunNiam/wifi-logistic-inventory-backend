package route

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/controller"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/repository"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func StockInRoute(router *gin.Engine, db *gorm.DB, redisClient *redis.Client, validate *validator.Validate) {
	stockInService := service.NewStockInService(
		repository.NewStockInRepository(),
		repository.NewProductRepository(),
		db,
		redisClient,
		validate,
	)
	stockInController := controller.NewStockInController(stockInService)

	router.GET("/stock-in", auth.Auth(stockInController.FindAll, []string{}))
	router.GET("/stock-in/:id", auth.Auth(stockInController.FindByID, []string{}))
	router.POST("/stock-in", auth.Auth(stockInController.Create, []string{}))
	router.PUT("/stock-in/:id", auth.Auth(stockInController.Update, []string{}))
	router.DELETE("/stock-in/:id", auth.Auth(stockInController.Delete, []string{}))
}
