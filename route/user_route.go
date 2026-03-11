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

func UserRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	userService := service.NewUserService(
		repository.NewUserRepository(),
		db,
		validate)
	userController := controller.NewUserController(userService)

	// Public routes
	router.POST("/users/login", userController.Login)

	// Protected routes
	router.GET("/users", auth.Auth(userController.FindAll, []string{}))
	router.GET("/users/:id", auth.Auth(userController.FindByID, []string{}))
	router.POST("/users", auth.Auth(userController.Create, []string{}))
	router.PUT("/users/:id", auth.Auth(userController.Update, []string{}))
	router.DELETE("/users/:id", auth.Auth(userController.Delete, []string{}))
	router.PUT("/users/:id/password", auth.Auth(userController.ChangePassword, []string{}))
}
