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
	"gorm.io/gorm"
)

type SupplierServiceImpl struct {
	SupplierRepository repository.SupplierRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewSupplierService(
	supplierRepository repository.SupplierRepository,
	db *gorm.DB,
	validate *validator.Validate,
) SupplierService {
	return &SupplierServiceImpl{
		SupplierRepository: supplierRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *SupplierServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.SupplierResponse {
	suppliers := service.SupplierRepository.FindAll(service.DB, filters)
	return suppliers.ToSupplierResponses()
}

func (service *SupplierServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.SupplierResponse {
	supplier := service.SupplierRepository.FindByID(service.DB, id)
	return supplier.ToSupplierResponse()
}

func (service *SupplierServiceImpl) Create(auth *auth.AccessDetails, request *web.SupplierCreateRequest, c *gin.Context) web.SupplierResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	supplier := domain.Supplier{
		Code:        request.Code,
		Name:        request.Name,
		Contact:     request.Contact,
		Phone:       request.Phone,
		Email:       request.Email,
		Address:     request.Address,
		Status:      request.Status,
		CreatedByID: auth.ID,
	}

	createdSupplier, err := service.SupplierRepository.Create(service.DB, &supplier)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	return createdSupplier.ToSupplierResponse()
}

func (service *SupplierServiceImpl) Update(auth *auth.AccessDetails, id uint, request *web.SupplierUpdateRequest, c *gin.Context) web.SupplierResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	supplier := domain.Supplier{
		Model:       gorm.Model{ID: id},
		Code:        request.Code,
		Name:        request.Name,
		Contact:     request.Contact,
		Phone:       request.Phone,
		Email:       request.Email,
		Address:     request.Address,
		Status:      request.Status,
		UpdatedByID: auth.ID,
	}

	updatedSupplier := service.SupplierRepository.Update(service.DB, &supplier)

	return updatedSupplier.ToSupplierResponse()
}

func (service *SupplierServiceImpl) Delete(auth *auth.AccessDetails, id uint, c *gin.Context) {
	service.SupplierRepository.Delete(service.DB, id, auth.ID)
}
