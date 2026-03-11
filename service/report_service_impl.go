package service

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/auth"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/helper"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/domain"
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReportServiceImpl struct {
	DB *gorm.DB
}

func NewReportService(db *gorm.DB) ReportService {
	return &ReportServiceImpl{
		DB: db,
	}
}

func (service *ReportServiceImpl) GenerateStockSummary(auth *auth.AccessDetails, startDate, endDate string, c *gin.Context) web.StockSummaryReport {
	var totalProducts int64
	var totalStock int64
	var totalStockValue float64
	var lowStockCount int64

	// Count total products
	service.DB.Model(&domain.Product{}).Count(&totalProducts)

	// Sum total stock
	service.DB.Model(&domain.Product{}).Select("COALESCE(SUM(stock), 0)").Scan(&totalStock)

	// Sum total stock value
	service.DB.Model(&domain.Product{}).Select("COALESCE(SUM(stock * price), 0)").Scan(&totalStockValue)

	// Count low stock
	service.DB.Model(&domain.Product{}).Where("stock <= min_stock").Count(&lowStockCount)

	// Get stock in count for period
	var stockInCount int64
	var stockInQty int64
	tx := service.DB.Model(&domain.StockIn{})
	if startDate != "" {
		tx = tx.Where("date >= ?", startDate)
	}
	if endDate != "" {
		tx = tx.Where("date <= ?", endDate)
	}
	tx.Count(&stockInCount)
	tx.Select("COALESCE(SUM(quantity), 0)").Scan(&stockInQty)

	// Get stock out count for period
	var stockOutCount int64
	var stockOutQty int64
	tx = service.DB.Model(&domain.StockOut{})
	if startDate != "" {
		tx = tx.Where("date >= ?", startDate)
	}
	if endDate != "" {
		tx = tx.Where("date <= ?", endDate)
	}
	tx.Count(&stockOutCount)
	tx.Select("COALESCE(SUM(quantity), 0)").Scan(&stockOutQty)

	return web.StockSummaryReport{
		TotalProducts:   totalProducts,
		TotalStock:      totalStock,
		TotalStockValue: totalStockValue,
		LowStockCount:   lowStockCount,
		StockInCount:    stockInCount,
		StockInQty:      stockInQty,
		StockOutCount:   stockOutCount,
		StockOutQty:     stockOutQty,
		StartDate:       startDate,
		EndDate:         endDate,
	}
}

func (service *ReportServiceImpl) GenerateStockIn(auth *auth.AccessDetails, startDate, endDate string, c *gin.Context) []web.StockInResponse {
	var stockIns domain.StockIns

	tx := service.DB.Model(&domain.StockIn{}).Preload("Product").Preload("Supplier").Preload("User")

	if startDate != "" {
		tx = tx.Where("date >= ?", startDate)
	}
	if endDate != "" {
		tx = tx.Where("date <= ?", endDate)
	}

	err := tx.Order("date DESC").Find(&stockIns).Error
	helper.PanicIfError(err)

	return stockIns.ToStockInResponses()
}

func (service *ReportServiceImpl) GenerateStockOut(auth *auth.AccessDetails, startDate, endDate string, c *gin.Context) []web.StockOutResponse {
	var stockOuts domain.StockOuts

	tx := service.DB.Model(&domain.StockOut{}).Preload("Product").Preload("User")

	if startDate != "" {
		tx = tx.Where("date >= ?", startDate)
	}
	if endDate != "" {
		tx = tx.Where("date <= ?", endDate)
	}

	err := tx.Order("date DESC").Find(&stockOuts).Error
	helper.PanicIfError(err)

	return stockOuts.ToStockOutResponses()
}

func (service *ReportServiceImpl) GenerateLowStock(auth *auth.AccessDetails, c *gin.Context) []web.LowStockProduct {
	var products []domain.Product

	err := service.DB.Model(&domain.Product{}).Where("stock <= min_stock").Find(&products).Error
	helper.PanicIfError(err)

	lowStockItems := make([]web.LowStockProduct, len(products))
	for i, product := range products {
		lowStockItems[i] = web.LowStockProduct{
			ID:       product.ID,
			Code:     product.Code,
			Name:     product.Name,
			Stock:    product.Stock,
			MinStock: product.MinStock,
			Unit:     product.Unit,
		}
	}

	return lowStockItems
}
