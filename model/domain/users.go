package domain

import (
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
	FullName    string `gorm:"size:200;uniqueIndex:idx_users"`
	LegalName   string `gorm:"size:200"`
	Password    string `gorm:"size:100"`
	Role        string `gorm:"size:100"`
	NumberPhone string `gorm:"size:15;uniqueIndex:idx_users"`
	Email       string `gorm:"size:100;uniqueIndex:idx_users"`
	Status      string `gorm:"size:20;default:'active'"`
}

func (user *User) ToUserResponse() web.UserResponse {
	return web.UserResponse{
		ID:          user.ID,
		FullName:    user.FullName,
		LegalName:   user.LegalName,
		NumberPhone: user.NumberPhone,
		Email:       user.Email,
		Role:        user.Role,
		Status:      user.Status,
	}
}

func (users Users) ToUserResponses() []web.UserResponse {
	userResponses := []web.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, user.ToUserResponse())
	}
	return userResponses
}
