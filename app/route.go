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

	// CORS middleware for frontend
	router.Use(CORSMiddleware())

	// route path
	route.UserRoute(router, db, validate)
	route.ProductRoute(router, db, redisClient, validate)
	route.SupplierRoute(router, db, redisClient, validate)
	route.CategoryRoute(router, db, redisClient, validate)
	route.StockInRoute(router, db, redisClient, validate)
	route.StockOutRoute(router, db, redisClient, validate)
	route.DashboardRoute(router, db, redisClient)
	route.ReportRoute(router, db)
	// route.TransactionRoute(router, db, validate)

	return router
}

// CORSMiddleware adds CORS headers for frontend integration
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
