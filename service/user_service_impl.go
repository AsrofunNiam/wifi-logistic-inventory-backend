package service

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/exception"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/repository"
	"github.com/gin-gonic/gin"
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

func (service *UserServiceImpl) FindAll(authDetails *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.UserResponse {
	users := service.UserRepository.FindAll(service.DB, filters)
	return users.ToUserResponses()
}

func (service *UserServiceImpl) FindByID(authDetails *auth.AccessDetails, id *uint, c *gin.Context) web.UserResponse {
	user := service.UserRepository.FindByID(service.DB, id)
	return user.ToUserResponse()
}

func (service *UserServiceImpl) Create(authDetails *auth.AccessDetails, request *web.UserCreateRequest, c *gin.Context) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user := domain.User{
		FullName:    request.FullName,
		LegalName:   request.LegalName,
		Email:       request.Email,
		Password:    string(hashedPassword),
		Role:        request.Role,
		CreatedByID: authDetails.ID,
	}

	createdUser, err := service.UserRepository.Create(service.DB, &user)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	return createdUser.ToUserResponse()
}

func (service *UserServiceImpl) Update(authDetails *auth.AccessDetails, id uint, request *web.UserUpdateRequest, c *gin.Context) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	user := domain.User{
		Model:       gorm.Model{ID: id},
		FullName:    request.FullName,
		LegalName:   request.LegalName,
		Email:       request.Email,
		Role:        request.Role,
		UpdatedByID: authDetails.ID,
	}

	updatedUser := service.UserRepository.Update(service.DB, &user)
	return updatedUser.ToUserResponse()
}

func (service *UserServiceImpl) Delete(authDetails *auth.AccessDetails, id uint, c *gin.Context) {
	service.UserRepository.Delete(service.DB, id, authDetails.ID)
}

func (service *UserServiceImpl) ChangePassword(authDetails *auth.AccessDetails, id uint, request *web.ChangePasswordRequest, c *gin.Context) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Get current user
	user := service.UserRepository.FindByID(service.DB, &id)

	// Verify old password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if err != nil {
		panic(exception.NewBadRequestError("Invalid old password"))
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	// Update password
	user.Password = string(hashedPassword)
	service.UserRepository.Update(service.DB, &user)
}
