package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
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

type ProductServiceImpl struct {
	ProductRepository     repository.ProductRepository
	TransactionRepository repository.TransactionRepository
	BalanceRepository     repository.BalanceRepository
	DB                    *gorm.DB
	RedisClient           *redis.Client
	Validate              *validator.Validate
}

func NewProductService(
	productRepository repository.ProductRepository,
	transactionRepository repository.TransactionRepository,
	limitRepository repository.BalanceRepository,
	db *gorm.DB,
	redisClient *redis.Client,
	validate *validator.Validate,
) ProductService {
	return &ProductServiceImpl{
		ProductRepository:     productRepository,
		TransactionRepository: transactionRepository,
		BalanceRepository:     limitRepository,
		DB:                    db,
		RedisClient:           redisClient,
		Validate:              validate,
	}
}

func (service *ProductServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.ProductResponse {
	ctx := context.Background()
	key := "products:all"

	// Cek cache di Redis
	data, err := service.RedisClient.Get(ctx, key).Result()
	if err == nil {
		// If cache products found
		var cachedProducts []web.ProductResponse
		if err := json.Unmarshal([]byte(data), &cachedProducts); err == nil {
			return cachedProducts
		}
	}

	// If cache not found
	products := service.ProductRepository.FindAll(service.DB, filters)
	productResponses := products.ToProductResponses() // Convert ke response DTO

	// Save products to Redis
	jsonData, err := json.Marshal(productResponses)
	if err == nil {
		_ = service.RedisClient.Set(ctx, key, jsonData, 100*time.Minute).Err()
	}

	return productResponses
}

func (service *ProductServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.ProductResponse {
	ctx := context.Background()
	key := fmt.Sprintf("product: %s", fmt.Sprintf("%d", *id))

	// Cek cache di Redis
	data, err := service.RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		// If cache product not found so get from database
		product := service.ProductRepository.FindByID(service.DB, id)
		productResponse := product.ToProductResponse()

		// Save to Redis
		jsonData, _ := json.Marshal(productResponse)
		_ = service.RedisClient.Set(ctx, key, jsonData, 100*time.Minute).Err()

		return productResponse
	} else if err != nil {
		helper.PanicIfError(err)
	}

	var cachedProduct web.ProductResponse
	_ = json.Unmarshal([]byte(data), &cachedProduct)
	return cachedProduct

}

func (service *ProductServiceImpl) FindImage(auth *auth.AccessDetails, imagesName string, c *gin.Context) []byte {
	// Open file
	fileProofPhoto := helper.PathToProduct + imagesName
	fileOpen, err := os.Open(fileProofPhoto)
	helper.PanicIfError(err)

	fileBytes, err := io.ReadAll(fileOpen)
	helper.PanicIfError(err)

	return fileBytes
}
func (service *ProductServiceImpl) Create(auth *auth.AccessDetails, request *web.ProductCreateRequest, c *gin.Context) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	err = os.MkdirAll(helper.PathToProduct, os.ModePerm)
	helper.PanicIfError(err)

	// create new file image
	var imageName string
	if len(request.ImageFile) > 0 {
		imageName = request.ImageFile[0].Filename

		filePath := helper.PathToProduct + imageName
		err = c.SaveUploadedFile(request.ImageFile[0], filePath)
		if err != nil {
			helper.PanicIfError(err)
		}
	}

	newProduct := &domain.Product{
		Name:        request.Name,
		Type:        request.Type,
		CompanyCode: request.CompanyCode,
		Description: request.Description,
		Images:      imageName,
		Available:   request.Available,
		CreatedByID: auth.ID,
	}

	newProduct, err = service.ProductRepository.Create(tx, newProduct)

	// If failed insert to database, delete image file
	if err != nil {
		_ = os.Remove(helper.PathToProduct + imageName)
		helper.PanicIfError(err)
	}

	return newProduct.ToProductResponse()
}

func (service *ProductServiceImpl) Update(auth *auth.AccessDetails, id uint, request *web.ProductUpdateRequest, c *gin.Context) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// Create folder if not exist
	err = os.MkdirAll(helper.PathToProduct, os.ModePerm)
	helper.PanicIfError(err)

	product := &domain.Product{
		Model:       gorm.Model{ID: uint(id)},
		UpdatedByID: auth.ID,
		Name:        request.Name,
		Type:        request.Type,
		Description: request.Description,
		Available:   request.Available,
	}

	// Validate images
	newImage := request.ImageFile[0].Filename
	oldImages := request.ImageName
	if len(request.ImageFile) > 0 {
		product.Images = newImage

		// Save new image
		err = c.SaveUploadedFile(request.ImageFile[0], helper.PathToProduct+newImage)
		helper.PanicIfError(err)

		// Delete image old if exists
		if oldImages != newImage {
			oldImagePath := helper.PathToProduct + oldImages
			err = os.Remove(oldImagePath)
			if err != nil && !os.IsNotExist(err) {
				helper.PanicIfError(err)
			}
		}
	} else {
		product.Images = oldImages
	}

	// Update product
	product = service.ProductRepository.Update(tx, product)

	return product.ToProductResponse()
}

func (service *ProductServiceImpl) Delete(auth *auth.AccessDetails, id uint, c *gin.Context) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	service.ProductRepository.Delete(tx, id, auth.ID)
}

// Group tRansaction Product
func (service *ProductServiceImpl) FindAllTransaction(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.TransactionResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	transaction := service.TransactionRepository.FindAll(tx, filters)
	return transaction.ToTransactionResponses()
}

func (service *ProductServiceImpl) CreateTransaction(auth *auth.AccessDetails, request *web.TransactionCreateRequest, c *gin.Context) web.TransactionResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	//  Validate request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Validate user role
	if auth.Role != "customer" {
		err := &exception.ErrorSendToResponse{Err: "Only customers can create loans"}
		helper.PanicIfError(err)
	}

	// Find product
	product := service.ProductRepository.FindByID(tx, &request.ProductID)

	// Find limit
	limit := service.BalanceRepository.FindByID(tx, &auth.ID)

	// Validate limit existence
	if limit.ID == 0 {
		err := &exception.ErrorSendToResponse{Err: "Balance not found"}
		helper.PanicIfError(err)
	}

	totalValue := product.ProductPrice.Price + request.AdminFee

	// Validate limit against instalment amount
	if limit.Value < totalValue {
		err := &exception.ErrorSendToResponse{Err: "Insufficient limit for the requested product"}
		helper.PanicIfError(err)
	}

	transaction := &domain.Transaction{
		// Required Fields
		CreatedByID: auth.ID,
		UserID:      auth.ID,

		// Fields for Loan
		OnTheRoad:       product.ProductPrice.Price,
		AdminFee:        request.AdminFee,
		TotalValue:      totalValue,
		ProductID:       request.ProductID,
		ProductName:     product.Name,
		TransactionType: request.TransactionType,
		NumberContract:  fmt.Sprintf("%d-%d-%d", auth.ID, product.ID, time.Now().Unix()),
	}

	transaction = service.TransactionRepository.Create(tx, transaction)

	// Update limit
	limit.Value -= totalValue
	service.BalanceRepository.Update(tx, &limit)

	return transaction.ToTransactionResponse()
}
