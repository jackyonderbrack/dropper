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
	// Załaduj zmienne środowiskowe z pliku .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Nie udało się załadować pliku .env, używane są domyślne wartości")
	}

	// Pobierz dane konfiguracyjne z pliku .env lub użyj domyślnych
	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		apiPort = "8080" // Domyślny port
	}

	// Połączenie z bazą danych
	database.ConnectDB()

	// Stworzenie routera
	r := mux.NewRouter()

	// Serwowanie plików statycznych
	// Gdybyśmy zostawili wyżej HandleFunc ze ścieżką ("/") to ówczesna
	// ścieżka nie zostałaby już uwzględniona.
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./client/dist"))))

	// Obsługa React Router (przekierowanie do index.html dla każdej nieznanej trasy)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./client/dist/index.html")
	})

	// Uruchamiamy serwer
	log.Printf("Serwer nasłuchuje na porcie: %s", apiPort)
	log.Fatal(http.ListenAndServe(":"+apiPort, r))
}
