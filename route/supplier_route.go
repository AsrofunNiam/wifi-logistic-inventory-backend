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

func SupplierRoute(router *gin.Engine, db *gorm.DB, redisClient *redis.Client, validate *validator.Validate) {
	supplierService := service.NewSupplierService(
		repository.NewSupplierRepository(),
		db,
		redisClient,
		validate,
	)
	supplierController := controller.NewSupplierController(supplierService)

	router.GET("/suppliers", auth.Auth(supplierController.FindAll, []string{}))
	router.GET("/suppliers/:id", auth.Auth(supplierController.FindByID, []string{}))
	router.POST("/suppliers", auth.Auth(supplierController.Create, []string{}))
	router.PUT("/suppliers/:id", auth.Auth(supplierController.Update, []string{}))
	router.DELETE("/suppliers/:id", auth.Auth(supplierController.Delete, []string{}))
}
