package controller

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(context *gin.Context)
	FindAll(context *gin.Context, auth *auth.AccessDetails)
	FindByID(context *gin.Context, auth *auth.AccessDetails)
	Create(context *gin.Context, auth *auth.AccessDetails)
	Update(context *gin.Context, auth *auth.AccessDetails)
	Delete(context *gin.Context, auth *auth.AccessDetails)
	ChangePassword(context *gin.Context, auth *auth.AccessDetails)
}
