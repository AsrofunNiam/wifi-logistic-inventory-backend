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

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewCategoryService(
	categoryRepository repository.CategoryRepository,
	db *gorm.DB,
	validate *validator.Validate,
) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.CategoryResponse {
	categories := service.CategoryRepository.FindAll(service.DB, filters)
	return categories.ToCategoryResponses()
}

func (service *CategoryServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.CategoryResponse {
	category := service.CategoryRepository.FindByID(service.DB, id)
	return category.ToCategoryResponse()
}

func (service *CategoryServiceImpl) Create(auth *auth.AccessDetails, request *web.CategoryCreateRequest, c *gin.Context) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	category := domain.Category{
		Name:        request.Name,
		Description: request.Description,
		CreatedByID: auth.ID,
	}

	createdCategory, err := service.CategoryRepository.Create(service.DB, &category)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	return createdCategory.ToCategoryResponse()
}

func (service *CategoryServiceImpl) Update(auth *auth.AccessDetails, id uint, request *web.CategoryUpdateRequest, c *gin.Context) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	category := domain.Category{
		Model:       gorm.Model{ID: id},
		Name:        request.Name,
		Description: request.Description,
		UpdatedByID: auth.ID,
	}

	updatedCategory := service.CategoryRepository.Update(service.DB, &category)

	return updatedCategory.ToCategoryResponse()
}

func (service *CategoryServiceImpl) Delete(auth *auth.AccessDetails, id uint, c *gin.Context) {
	service.CategoryRepository.Delete(service.DB, id, auth.ID)
}
