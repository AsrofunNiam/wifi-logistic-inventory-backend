package app

import (
	"fmt"
	"runtime/debug"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/exception"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/route"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ErrorHandler
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
				exception.ErrorHandler(c, err)
			}
		}()
		c.Next()
	}
}

func NewRouter(db *gorm.DB, redisClient *redis.Client, validate *validator.Validate) *gin.Engine {

	router := gin.New()

	//  exception middleware
	router.Use(ErrorHandler())
	router.UseRawPath = true

	// route path
	route.UserRoute(router, db, validate)
	route.ProductRoute(router, db, redisClient, validate)
	// route.TransactionRoute(router, db, validate)

	return router
}
