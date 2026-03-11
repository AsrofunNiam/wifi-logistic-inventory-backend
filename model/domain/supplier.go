package domain

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"gorm.io/gorm"
)

type Suppliers []Supplier
type Supplier struct {
	gorm.Model
	CreatedByID uint  `gorm:"default:null"`
	UpdatedByID uint  `gorm:"default:null"`
	DeletedByID *uint `gorm:"default:null"`

	// Required Fields
	Code    string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Name    string `gorm:"type:varchar(255);not null"`
	Contact string `gorm:"type:varchar(200)"`
	Phone   string `gorm:"type:varchar(20)"`
	Email   string `gorm:"type:varchar(100)"`
	Address string `gorm:"type:text"`
	Status  string `gorm:"type:varchar(20);default:'active'"` // active, inactive

	// Relations
	Products []Product `gorm:"foreignKey:SupplierID"`
}

func (supplier *Supplier) ToSupplierResponse() web.SupplierResponse {
	return web.SupplierResponse{
		ID:        supplier.ID,
		Code:      supplier.Code,
		Name:      supplier.Name,
		Contact:   supplier.Contact,
		Phone:     supplier.Phone,
		Email:     supplier.Email,
		Address:   supplier.Address,
		Status:    supplier.Status,
		CreatedAt: supplier.CreatedAt,
		UpdatedAt: supplier.UpdatedAt,
	}
}

func (suppliers Suppliers) ToSupplierResponses() []web.SupplierResponse {
	supplierResponses := []web.SupplierResponse{}
	for _, supplier := range suppliers {
		supplierResponses = append(supplierResponses, supplier.ToSupplierResponse())
	}
	return supplierResponses
}
