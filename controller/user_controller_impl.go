package controller

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/service"
	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Login(c *gin.Context) {
	userAgent := c.GetHeader("User-Agent")
	remoteAddress := c.Request.RemoteAddr
	request := web.UserLoginRequest{}
	helper.ReadFromRequestBody(c, &request)

	// Decode access_token
	decodedToken, err := base64.StdEncoding.DecodeString(request.AccessTokenLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Success: false,
			Message: "Invalid access token",
		})
		return
	}

	// split password and name
	credentials := strings.Split(string(decodedToken), ":")
	if len(credentials) != 2 {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Success: false,
			Message: "Invalid credentials format",
		})
		return
	}
	identity := credentials[0]
	password := credentials[1]

	tokenResponse := controller.UserService.Login(&identity, &password, &userAgent, &remoteAddress, &request)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Login success",
		Data:    tokenResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}
