package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type ApiResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	dbHost := getEnv("DB_HOST", "postgres-service") //nome service db
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "mydb")
	dbUser := getEnv("DB_USER", "myuser")
	dbPassword := getEnv("DB_PASSWORD", "mypassword")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Errore apertura connessione DB:", err)
	}
	defer db.Close()

	http.HandleFunc("/api/check-db", func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, "Metodo non consentito", http.StatusMethodNotAllowed)
			return
		}

		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ApiResponse{
				Status:  "error",
				Message: "DB non raggiungibile",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ApiResponse{
			Status:  "ok",
			Message: "Backend Go OK - DB raggiungibile",
		})
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("Backend in ascolto sulla porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
