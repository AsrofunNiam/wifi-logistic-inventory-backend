package controller

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
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

func (controller *UserControllerImpl) FindAll(c *gin.Context, authDetails *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "full_name.like", "email.like", "role.eq")
	userResponses := controller.UserService.FindAll(authDetails, &filters, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(userResponses),
		Data:    userResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) FindByID(c *gin.Context, authDetails *auth.AccessDetails) {
	userIDParam := c.Param("id")
	userID := helper.StringToUint(userIDParam)

	userResponse := controller.UserService.FindByID(authDetails, &userID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(userResponse),
		Data:    userResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Create(c *gin.Context, authDetails *auth.AccessDetails) {
	var request web.UserCreateRequest
	err := c.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	userResponse := controller.UserService.Create(authDetails, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "User created successfully",
		Data:    userResponse,
	}
	c.JSON(http.StatusCreated, webResponse)
}

func (controller *UserControllerImpl) Update(c *gin.Context, authDetails *auth.AccessDetails) {
	userIDParam := c.Param("id")
	userID := helper.StringToUint(userIDParam)

	var request web.UserUpdateRequest
	err := c.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	userResponse := controller.UserService.Update(authDetails, userID, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    userResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Delete(c *gin.Context, authDetails *auth.AccessDetails) {
	userIDParam := c.Param("id")
	userID := helper.StringToUint(userIDParam)

	controller.UserService.Delete(authDetails, userID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "User deleted successfully",
		Data:    nil,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) ChangePassword(c *gin.Context, authDetails *auth.AccessDetails) {
	userIDParam := c.Param("id")
	userID := helper.StringToUint(userIDParam)

	var request web.ChangePasswordRequest
	err := c.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	controller.UserService.ChangePassword(authDetails, userID, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Password changed successfully",
		Data:    nil,
	}
	c.JSON(http.StatusOK, webResponse)
}
