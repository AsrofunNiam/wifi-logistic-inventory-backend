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

func ProductRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {
	productService := service.NewProductService(
		repository.NewProductRepository(),
		db,
		validate,
	)
	productController := controller.NewProductController(productService)
	router.GET("/products", auth.Auth(productController.FindAll, []string{}))
	router.GET("/products/:id", auth.Auth(productController.FindByID, []string{}))
	router.POST("/products", auth.Auth(productController.Create, []string{}))
	router.PUT("/products/:id", auth.Auth(productController.Update, []string{}))
	router.DELETE("/products/:id", auth.Auth(productController.Delete, []string{}))
}
