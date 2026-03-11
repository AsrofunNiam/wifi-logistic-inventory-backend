package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(context *gin.Context)
}
