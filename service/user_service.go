package service

import "github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"

type UserService interface {
	Login(identity, password, userAgent, remoteAddress *string, request *web.UserLoginRequest) web.TokenResponse
}
