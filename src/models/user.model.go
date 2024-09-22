package models

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null"`
	Email       string    `gorm:"size:255;unique;not null"`
	Password    string    `gorm:"size:255;not null"`
	IsAdmin     bool      `gorm:"default:false"`
	Company 	Company   `gorm:"embedded"`
	NIP         string    `gorm:"size:255"`
	Address     Address   `gorm:"embedded"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Model dla danych firmy (opcjonalne)
type Company struct {
	Name string `gorm:"size:255"`
	NIP  string `gorm:"size:100"`
}

// Model dla adresu (opcjonalne)
type Address struct {
	Street  string `gorm:"size:255"`
	City    string `gorm:"size:100"`
	ZipCode string `gorm:"size:50"`
	Country string `gorm:"size:100"`
}