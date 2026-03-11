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

func CategoryRoute(router *gin.Engine, db *gorm.DB, redisClient *redis.Client, validate *validator.Validate) {
	categoryService := service.NewCategoryService(
		repository.NewCategoryRepository(),
		db,
		redisClient,
		validate,
	)
	categoryController := controller.NewCategoryController(categoryService)

	router.GET("/categories", auth.Auth(categoryController.FindAll, []string{}))
	router.GET("/categories/:id", auth.Auth(categoryController.FindByID, []string{}))
	router.POST("/categories", auth.Auth(categoryController.Create, []string{}))
	router.PUT("/categories/:id", auth.Auth(categoryController.Update, []string{}))
	router.DELETE("/categories/:id", auth.Auth(categoryController.Delete, []string{}))
}
