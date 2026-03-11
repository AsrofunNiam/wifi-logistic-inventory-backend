package web

import "time"

type DashboardStatsResponse struct {
	TotalProducts    int64             `json:"total_products"`
	TotalSuppliers   int64             `json:"total_suppliers"`
	TotalStockIn     int64             `json:"total_stock_in"`
	TotalStockOut    int64             `json:"total_stock_out"`
	TotalStockInQty  int64             `json:"total_stock_in_qty"`
	TotalStockOutQty int64             `json:"total_stock_out_qty"`
	LowStockCount    int64             `json:"low_stock_count"`
	LowStockProducts []LowStockProduct `json:"low_stock_products"`
	RecentActivities []RecentActivity  `json:"recent_activities"`
}

type LowStockProduct struct {
	ID       uint   `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
	MinStock int    `json:"min_stock"`
	Unit     string `json:"unit"`
}

type RecentActivity struct {
	ID          uint      `json:"id"`
	Type        string    `json:"type"` // "stock_in" or "stock_out"
	Code        string    `json:"code"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Date        string    `json:"date"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type StockMovementChartResponse struct {
	Date     string `json:"date"`
	StockIn  int    `json:"stock_in"`
	StockOut int    `json:"stock_out"`
}
