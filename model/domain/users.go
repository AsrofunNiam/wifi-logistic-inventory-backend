package domain

import (
	"time"

	"github.com/AsrofunNiam/wifi-logistic-inventory-backend/model/web"
	"gorm.io/gorm"
)

type Users []User
type User struct {
	// Required Fields
	gorm.Model
	CreatedByID uint  `gorm:""`
	UpdatedByID uint  `gorm:""`
	DeletedByID *uint `gorm:""`

	// Fields
	FullName      string    `gorm:"size:200;uniqueIndex:idx_users"`
	LegalName     string    `gorm:"size:200"`
	PlaceOfBirth  string    `gorm:"size:200"`
	DateOfBirth   time.Time `gorm:"type:date;not null"`
	Salary        float64   `gorm:"type:decimal(20,2);not null"`
	IdentityImage string    `gorm:"size:200"`
	FaceImage     string    `gorm:"size:200"`
	Password      string    `gorm:"size:100"`
	Role          string    `gorm:"size:100"`
	NumberPhone   string    `gorm:"size:15;uniqueIndex:idx_users"`
	Email         string    `gorm:"size:100;uniqueIndex:idx_users"`
}

func (user *User) ToUserResponse() web.UserResponse {
	return web.UserResponse{
		// Required Fields
		ID: user.ID,
		// Fields
		FullName:    user.FullName,
		LegalName:   user.LegalName,
		NumberPhone: user.NumberPhone,
		Email:       user.Email,
		Role:        user.Role,
	}
}

func (users Users) ToUserResponses() []web.UserResponse {
	userResponses := []web.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, user.ToUserResponse())
	}
	return userResponses
}
