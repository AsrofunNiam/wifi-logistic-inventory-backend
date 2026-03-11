package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DashboardServiceImpl struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewDashboardService(
	db *gorm.DB,
	redisClient *redis.Client,
) DashboardService {
	return &DashboardServiceImpl{
		DB:          db,
		RedisClient: redisClient,
	}
}

func (service *DashboardServiceImpl) GetStats(auth *auth.AccessDetails, c *gin.Context) web.DashboardStatsResponse {
	ctx := context.Background()
	key := "dashboard:stats"

	// Check cache in Redis
	data, err := service.RedisClient.Get(ctx, key).Result()
	if err == nil {
		var cachedStats web.DashboardStatsResponse
		if err := json.Unmarshal([]byte(data), &cachedStats); err == nil {
			return cachedStats
		}
	}

	// Calculate stats from database
	var totalProducts int64
	var totalSuppliers int64
	var totalStockIn int64
	var totalStockOut int64
	var totalStockInQty int64
	var totalStockOutQty int64
	var lowStockCount int64

	// Count total products
	err = service.DB.Model(&domain.Product{}).Count(&totalProducts).Error
	helper.PanicIfError(err)

	// Count total suppliers
	err = service.DB.Model(&domain.Supplier{}).Count(&totalSuppliers).Error
	helper.PanicIfError(err)

	// Count today's stock in
	today := time.Now().Format("2006-01-02")
	err = service.DB.Model(&domain.StockIn{}).Where("DATE(date) = ?", today).Count(&totalStockIn).Error
	helper.PanicIfError(err)

	// Sum today's stock in quantity
	service.DB.Model(&domain.StockIn{}).Where("DATE(date) = ?", today).Select("COALESCE(SUM(quantity), 0)").Scan(&totalStockInQty)

	// Count today's stock out
	err = service.DB.Model(&domain.StockOut{}).Where("DATE(date) = ?", today).Count(&totalStockOut).Error
	helper.PanicIfError(err)

	// Sum today's stock out quantity
	service.DB.Model(&domain.StockOut{}).Where("DATE(date) = ?", today).Select("COALESCE(SUM(quantity), 0)").Scan(&totalStockOutQty)

	// Count low stock products
	err = service.DB.Model(&domain.Product{}).Where("stock <= min_stock").Count(&lowStockCount).Error
	helper.PanicIfError(err)

	// Get low stock products
	var lowStockProducts []domain.Product
	err = service.DB.Model(&domain.Product{}).Where("stock <= min_stock").Preload("Category").Limit(10).Find(&lowStockProducts).Error
	helper.PanicIfError(err)

	lowStockItems := make([]web.LowStockProduct, len(lowStockProducts))
	for i, product := range lowStockProducts {
		lowStockItems[i] = web.LowStockProduct{
			ID:       product.ID,
			Code:     product.Code,
			Name:     product.Name,
			Stock:    product.Stock,
			MinStock: product.MinStock,
			Unit:     product.Unit,
		}
	}

	// Get recent activities
	var recentStockIns []domain.StockIn
	var recentStockOuts []domain.StockOut

	service.DB.Model(&domain.StockIn{}).
		Preload("Product").
		Preload("Supplier").
		Preload("User").
		Order("created_at DESC").
		Limit(5).
		Find(&recentStockIns)

	service.DB.Model(&domain.StockOut{}).
		Preload("Product").
		Preload("User").
		Order("created_at DESC").
		Limit(5).
		Find(&recentStockOuts)

	// Combine and sort activities
	activities := make([]web.RecentActivity, 0, len(recentStockIns)+len(recentStockOuts))

	for _, stockIn := range recentStockIns {
		activities = append(activities, web.RecentActivity{
			ID:          stockIn.ID,
			Type:        "stock_in",
			Code:        stockIn.Code,
			ProductName: stockIn.Product.Name,
			Quantity:    stockIn.Quantity,
			Date:        stockIn.Date.Format("2006-01-02"),
			CreatedBy:   stockIn.User.FullName,
			CreatedAt:   stockIn.CreatedAt,
		})
	}

	for _, stockOut := range recentStockOuts {
		activities = append(activities, web.RecentActivity{
			ID:          stockOut.ID,
			Type:        "stock_out",
			Code:        stockOut.Code,
			ProductName: stockOut.Product.Name,
			Quantity:    stockOut.Quantity,
			Date:        stockOut.Date.Format("2006-01-02"),
			CreatedBy:   stockOut.User.FullName,
			CreatedAt:   stockOut.CreatedAt,
		})
	}

	// Sort by created_at descending and limit to 10
	for i := 0; i < len(activities)-1; i++ {
		for j := i + 1; j < len(activities); j++ {
			if activities[j].CreatedAt.After(activities[i].CreatedAt) {
				activities[i], activities[j] = activities[j], activities[i]
			}
		}
	}

	if len(activities) > 10 {
		activities = activities[:10]
	}

	stats := web.DashboardStatsResponse{
		TotalProducts:    totalProducts,
		TotalSuppliers:   totalSuppliers,
		TotalStockIn:     totalStockIn,
		TotalStockOut:    totalStockOut,
		TotalStockInQty:  totalStockInQty,
		TotalStockOutQty: totalStockOutQty,
		LowStockCount:    lowStockCount,
		LowStockProducts: lowStockItems,
		RecentActivities: activities,
	}

	// Save to Redis
	jsonData, err := json.Marshal(stats)
	if err == nil {
		_ = service.RedisClient.Set(ctx, key, jsonData, 5*time.Minute).Err()
	}

	return stats
}
