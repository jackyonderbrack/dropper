# Dropper - sklep Example

## Wersja 1.0.0

Dropper to aplikacja backendowa napisana w Go, która serwuje aplikację frontendową zbudowaną przy użyciu Vite, React i TypeScript. Projekt ten obejmuje zarówno część klienta, jak i panel administracyjny.
Stworzył - Michał Łuczak [ MILU CONNECT ]

## Struktura projektu

```
dropper/
├── src/
│   ├── database/
│   │   └── config.go
│   ├── models/
│   │   └── user.model.go
│   ├── utils/
│   │   └── hash.go
│   ├── docs/
│   └── tutorial.md
├── client/
│   ├── src/
│   │   └── **
│   └── dist/ (folder zbudowany przez Vite)
├── \\.env
├── go.mod
└── go.sum

```

## Wymagania

Aby uruchomić ten projekt, będziesz potrzebować:

-   Go 1.18 lub nowszy
-   Node.js 20+ i npm (do budowy aplikacji Vite) # lub yarn
-   MySQL
-   Git

## Instalacja

1. Sklonuj repozytorium:

    ```bash
    git clone https://gitlab.com/twoje_repozytorium/dropper.git
    cd dropper
    ```

2. Backend (Go):

    - Zainstaluj zależności Go:

        ```bash
        cd backend
        go mod download
        ```

    - Skonfiguruj plik `.env` z ustawieniami bazy danych, jeśli używasz bazy danych.

3. Frontend (Vite + React + TypeScript):

    - Przejdź do katalogu klienta i zainstaluj zależności:

        ```bash
        cd client
        npm install  # lub yarn install
        ```

4. Zbuduj aplikację frontendową:

    ```bash
    npm run build  # lub yarn build
    ```

    Pliki wynikowe będą znajdować się w `client/dist/`.

## Uruchomienie

### Backend (Go):

Uruchom backend z poziomu katalogu `backend`:

```bash
go run src/main.go
```

### Frontend (Vite):

Po zbudowaniu aplikacji, frontend zostanie serwowany przez backend. Możesz uzyskać dostęp do aplikacji w przeglądarce pod adresem `http://localhost:<PORT>`.

## Testowanie

Aby przetestować aplikację lokalnie:

1. Uruchom serwer backendowy.
2. Otwórz przeglądarkę i przejdź do `http://localhost:<PORT>`, aby zobaczyć aplikację frontendową.

## Kontrybucja

Jeśli chcesz przyczynić się do rozwoju projektu, zrób fork repozytorium, utwórz swoją gałąź (`git checkout -b feature/nazwa-feature`), dodaj swoje zmiany i stwórz pull request.

## Licencja

Projekt jest licencjonowany na zasadach licencji MIT. Zobacz plik [LICENSE](LICENSE) po więcej informacji.
