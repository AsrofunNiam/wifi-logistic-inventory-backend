package service

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/exception"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/repository"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(
	userRepository repository.UserRepository,
	db *gorm.DB,
	validate *validator.Validate,
) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Login(identity, password, userAgent, remoteAddress *string, request *web.UserLoginRequest) web.TokenResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// validate
	err = service.Validate.Struct(request)
	helper.PanicIfError(err)

	user := service.UserRepository.Login(tx, identity)

	hashedPassword := []byte(*password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), hashedPassword)
	if err != nil {
		helper.PanicIfError(exception.ErrUnauthorized)
	}
	ts, err := auth.CreateToken(&user, userAgent, remoteAddress, func(userID uint, tokenDetails *auth.TokenDetails) {})
	if err != nil {
		helper.PanicIfError(err)
	}

	token := web.TokenResponse{
		ID:           user.ID,
		FullName:     user.FullName,
		LegalName:    user.LegalName,
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}

	return token
}
