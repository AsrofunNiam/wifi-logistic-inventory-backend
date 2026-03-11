package service

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Login(identity, password, userAgent, remoteAddress *string, request *web.UserLoginRequest) web.TokenResponse
	FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.UserResponse
	FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.UserResponse
	Create(auth *auth.AccessDetails, request *web.UserCreateRequest, c *gin.Context) web.UserResponse
	Update(auth *auth.AccessDetails, id uint, request *web.UserUpdateRequest, c *gin.Context) web.UserResponse
	Delete(auth *auth.AccessDetails, id uint, c *gin.Context)
	ChangePassword(auth *auth.AccessDetails, id uint, request *web.ChangePasswordRequest, c *gin.Context)
}
