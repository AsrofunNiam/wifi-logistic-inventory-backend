package web

type StockSummaryReport struct {
	TotalProducts   int64   `json:"total_products"`
	TotalStock      int64   `json:"total_stock"`
	TotalStockValue float64 `json:"total_stock_value"`
	LowStockCount   int64   `json:"low_stock_count"`
	StockInCount    int64   `json:"stock_in_count"`
	StockInQty      int64   `json:"stock_in_qty"`
	StockOutCount   int64   `json:"stock_out_count"`
	StockOutQty     int64   `json:"stock_out_qty"`
	StartDate       string  `json:"start_date"`
	EndDate         string  `json:"end_date"`
}

type StockMovementReport struct {
	Date     string `json:"date"`
	StockIn  int64  `json:"stock_in"`
	StockOut int64  `json:"stock_out"`
}
