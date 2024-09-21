package main

import (
	"backend/src/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	/*
	Załaduj zmienne środowiskowe z pliku .env (opcjonalne) 
	*/
    err := godotenv.Load()
    if err != nil {
        log.Println("Nie udało się załadować pliku .env, używane są domyślne wartości")
    }

    // Pobierz dane konfiguracyjne z pliku .env lub użyj domyślnych
    apiPort := os.Getenv("API_PORT")
    if apiPort == "" {
        apiPort = "No port"
    }

	/*
	Połączenie z bazą danych z pliku src/database/config.go
	*/
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ustawienie routera
	r := mux.NewRouter()

	// Testowanie trasy routera
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Witaj w API!"))
	})

	// Uruchommy serwer
	log.Printf("Serwer nasłuchuje na porcie: %s", apiPort)
	log.Fatal(http.ListenAndServe(":"+apiPort, r))
}