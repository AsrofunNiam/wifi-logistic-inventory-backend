package domain

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"gorm.io/gorm"
)

type Companies []Company
type Company struct {
	gorm.Model
	CreatedByID uint   `gorm:""`
	UpdatedByID uint   `gorm:""`
	DeletedByID uint   `gorm:""`
	Name        string `gorm:"size:200;uniqueIndex:idx_companies"`
	Description string `gorm:"type:text"`
	Address     string `gorm:"type:text"`
}

func (company *Company) ToCompanyResponse() web.CompanyResponse {
	return web.CompanyResponse{
		ID:          company.ID,
		Name:        company.Name,
		Description: company.Description,
		Address:     company.Address,
	}
}

func (users Companies) ToCompanyResponses() []web.CompanyResponse {
	companyResponses := []web.CompanyResponse{}
	for _, user := range users {
		companyResponses = append(companyResponses, user.ToCompanyResponse())
	}
	return companyResponses
}
