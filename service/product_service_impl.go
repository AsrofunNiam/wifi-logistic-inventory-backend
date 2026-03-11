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

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewProductService(
	productRepository repository.ProductRepository,
	db *gorm.DB,
	validate *validator.Validate,
) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.ProductResponse {
	products := service.ProductRepository.FindAll(service.DB, filters)
	return products.ToProductResponses()
}

func (service *ProductServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.ProductResponse {
	product := service.ProductRepository.FindByID(service.DB, id)
	return product.ToProductResponse()
}

func (service *ProductServiceImpl) Create(auth *auth.AccessDetails, request *web.ProductCreateRequest, c *gin.Context) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	newProduct := &domain.Product{
		Code:        request.Code,
		Name:        request.Name,
		CategoryID:  request.CategoryID,
		SupplierID:  request.SupplierID,
		Description: request.Description,
		Stock:       request.Stock,
		MinStock:    request.MinStock,
		Unit:        request.Unit,
		Price:       request.Price,
		CreatedByID: auth.ID,
	}

	newProduct, err = service.ProductRepository.Create(tx, newProduct)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	return newProduct.ToProductResponse()
}

func (service *ProductServiceImpl) Update(auth *auth.AccessDetails, id uint, request *web.ProductUpdateRequest, c *gin.Context) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	product := &domain.Product{
		Model:       gorm.Model{ID: id},
		Code:        request.Code,
		Name:        request.Name,
		CategoryID:  request.CategoryID,
		SupplierID:  request.SupplierID,
		Description: request.Description,
		Stock:       request.Stock,
		MinStock:    request.MinStock,
		Unit:        request.Unit,
		Price:       request.Price,
		UpdatedByID: auth.ID,
	}

	product = service.ProductRepository.Update(tx, product)

	return product.ToProductResponse()
}

func (service *ProductServiceImpl) Delete(auth *auth.AccessDetails, id uint, c *gin.Context) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	service.ProductRepository.Delete(tx, id, auth.ID)
}
