package domain

import (
	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"gorm.io/gorm"
)

type Categories []Category
type Category struct {
	gorm.Model
	CreatedByID uint  `gorm:"default:null"`
	UpdatedByID uint  `gorm:"default:null"`
	DeletedByID *uint `gorm:"default:null"`

	// Required Fields
	Name        string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Description string `gorm:"type:text"`
}

func (category *Category) ToCategoryResponse() web.CategoryResponse {
	return web.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func (categories Categories) ToCategoryResponses() []web.CategoryResponse {
	categoryResponses := []web.CategoryResponse{}
	for _, category := range categories {
		categoryResponses = append(categoryResponses, category.ToCategoryResponse())
	}
	return categoryResponses
}
