package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// ConnectDB - funkcja nawiązująca połączenie z bazą danych
func ConnectDB() (*sql.DB, error) {
    // Załaduj zmienne środowiskowe z pliku .env (opcjonalne)
    err := godotenv.Load()
    if err != nil {
        log.Println("Nie udało się załadować pliku .env, używane są domyślne wartości")
    }

    // Pobierz dane konfiguracyjne z pliku .env lub użyj domyślnych
    dbUser := os.Getenv("DB_USER")
    if dbUser == "" {
        dbUser = "No user"
    }

    dbPassword := os.Getenv("DB_PASSWORD")
    if dbPassword == "" {
        dbPassword = "No password"
    }

    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" {
        dbHost = "localhost"
    }

    dbPort := os.Getenv("DB_PORT")
    if dbPort == "" {
        dbPort = "3306"
    }

    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "dropper"
    }

    // Utwórz string połączenia
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    // Nawiąż połączenie z bazą danych
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Sprawdź połączenie
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    log.Println("Połączono z bazą danych!")
    return db, nil
}