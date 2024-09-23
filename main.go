package main

import (
	"backend/src/database"
	"backend/src/routes"
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

	routes.RegisterUserRoutes(r, database.DB)

	// Serwowanie plików statycznych - obsługuje pliki w ./client/dist
	fs := http.FileServer(http.Dir("./client/dist"))

	// Serwowanie plików statycznych, jeśli istnieją
	r.PathPrefix("/static/").Handler(fs)

	// Główna obsługa tras - przekierowywanie na index.html, jeśli nie ma fizycznych plików
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Sprawdź, czy istnieje fizyczny plik dla żądanej ścieżki
		_, err := os.Stat("./client/dist" + r.URL.Path)
		if os.IsNotExist(err) {
			// Jeśli plik nie istnieje, serwuj index.html, aby React Router przejął kontrolę
			http.ServeFile(w, r, "./client/dist/index.html")
			return
		}

		// W przeciwnym razie, serwuj plik normalnie (np. pliki JavaScript, CSS)
		fs.ServeHTTP(w, r)
	})

	// Uruchamiamy serwer
	log.Printf("Serwer nasłuchuje na porcie: %s", apiPort)
	log.Fatal(http.ListenAndServe(":"+apiPort, r))
}
