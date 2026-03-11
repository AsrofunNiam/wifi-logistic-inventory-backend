package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/exception"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *gorm.DB
	RedisClient        *redis.Client
	Validate           *validator.Validate
}

func NewCategoryService(
	categoryRepository repository.CategoryRepository,
	db *gorm.DB,
	redisClient *redis.Client,
	validate *validator.Validate,
) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		RedisClient:        redisClient,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.CategoryResponse {
	ctx := context.Background()
	key := "categories:all"

	// Check cache in Redis
	data, err := service.RedisClient.Get(ctx, key).Result()
	if err == nil {
		var cachedCategories []web.CategoryResponse
		if err := json.Unmarshal([]byte(data), &cachedCategories); err == nil {
			return cachedCategories
		}
	}

	// If cache not found
	categories := service.CategoryRepository.FindAll(service.DB, filters)
	categoryResponses := categories.ToCategoryResponses()

	// Save categories to Redis
	jsonData, err := json.Marshal(categoryResponses)
	if err == nil {
		_ = service.RedisClient.Set(ctx, key, jsonData, 100*time.Minute).Err()
	}

	return categoryResponses
}

func (service *CategoryServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.CategoryResponse {
	ctx := context.Background()
	key := fmt.Sprintf("category:%d", *id)

	// Check cache in Redis
	data, err := service.RedisClient.Get(ctx, key).Result()
	if err == nil {
		var cachedCategory web.CategoryResponse
		if err := json.Unmarshal([]byte(data), &cachedCategory); err == nil {
			return cachedCategory
		}
	}

	// If cache not found, get from database
	category := service.CategoryRepository.FindByID(service.DB, id)
	categoryResponse := category.ToCategoryResponse()

	// Save to Redis
	jsonData, _ := json.Marshal(categoryResponse)
	_ = service.RedisClient.Set(ctx, key, jsonData, 100*time.Minute).Err()

	return categoryResponse
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

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "categories:all").Err()

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

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "categories:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("category:%d", id)).Err()

	return updatedCategory.ToCategoryResponse()
}

func (service *CategoryServiceImpl) Delete(auth *auth.AccessDetails, id uint, c *gin.Context) {
	service.CategoryRepository.Delete(service.DB, id, auth.ID)

	// Invalidate cache
	ctx := context.Background()
	_ = service.RedisClient.Del(ctx, "categories:all").Err()
	_ = service.RedisClient.Del(ctx, fmt.Sprintf("category:%d", id)).Err()
}
