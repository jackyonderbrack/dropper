package database

import (
	"backend/src/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB - Globalna zmienna dla instancji bazy danych
var DB *gorm.DB

// ConnectDB - Funkcja do nawiązywania połączenia z bazą danych
func ConnectDB() {
	// Pobieranie zmiennych środowiskowych
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Tworzenie DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Nawiązywanie połączenia z bazą danych
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Nie udało się połączyć z bazą danych:", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Nie udało się wykonać migracji:", err)
	}

	log.Println("Połączono z bazą danych!")
}
