package services

import (
	"backend/src/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewuserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (us *UserService) IsAnyUser() (bool, error) {
	var user models.User
	result := us.DB.First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}