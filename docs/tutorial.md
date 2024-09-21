# Tworzenie serwera w Go

Poniżej znajduje się przykład prostego serwera HTTP w Go:

## 1. Tworzenie serwera w folderze `/backend`

Wchodzimy w docelowy folder w edytorze i w terminalu wpisujemy:

```bash
go mod init <nazwa_modułu>
```

## 2. Połączenie z bazą danych:

Zakładając, że mamy uruchomiony serwer SQL w
jakiejś przestrzeni, możemy połączyć do bazy nasz serwer.
W naszym przypadku będzie to baza uruchomiona w serwerze lokalnym SQL: http://localhost:3306
Tworzymy folder i plik `backend/database/config.go`

```go
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
        dbUser = "root"
    }

    dbPassword := os.Getenv("DB_PASSWORD")
    if dbPassword == "" {
        dbPassword = "yourpassword"
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
        dbName = "my_store"
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
```

```
backend/
├── src/
│   ├── database/
│   │   └── config.go
│   └── main.go
├── docs/
│   └── tutorial.md
├── bin/
├── .env
├── go.mod
└── go.sum

```
