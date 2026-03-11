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

func ProductRoute(router *gin.Engine, db *gorm.DB, redisClient *redis.Client, validate *validator.Validate) {
	Products := service.NewProductService(
		repository.NewProductRepository(),
		repository.NewTransactionRepository(),
		repository.NewBalanceRepository(),
		db,
		redisClient,
		validate,
	)
	productController := controller.NewProductController(Products)
	router.GET("/products", auth.Auth(productController.FindAll, []string{}))
	router.GET("/products/:id", auth.Auth(productController.FindByID, []string{}))
	router.GET("/products/photo/:image_name", auth.Auth(productController.FindImage, []string{}))
	router.POST("/products", auth.Auth(productController.Create, []string{}))
	router.DELETE("/products/:id", auth.Auth(productController.Delete, []string{}))
	router.PUT("/products/:id", auth.Auth(productController.Update, []string{}))

	// Group transaction
	router.GET("/products/transactions", auth.Auth(productController.FindAllProductTransaction, []string{}))
	router.POST("/products/transactions", auth.Auth(productController.CreateProductTransaction, []string{}))
}
