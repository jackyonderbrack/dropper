package models

import (
	"time"
)

type User struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Password    string     `json:"-"`          // Password będzie zahashowane, nie zwracaj go w odpowiedziach
	IsAdmin     bool       `json:"is_admin"`
	Company     *Company   `json:"company"`    // Opcjonalne pole z danymi firmy (null jeśli brak)
	Address     *Address   `json:"address"`    // Opcjonalne pole z adresem (null jeśli brak)
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Model dla danych firmy (opcjonalne)
type Company struct {
	Name string `json:"name"`
	NIP  string `json:"nip"`
}

// Model dla adresu (opcjonalne)
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}