package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Eyob49/expense-tracker-api/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db *database.Queries
}

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (cfg *apiConfig) handlerListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := cfg.db.ListCategories(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		log.Printf("Failed to fetch categories")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		log.Fatalf("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		log.Fatalf("Error connecting to the database: %v", err)
	}

	dbQueries := database.New(db)

	apiCfg := &apiConfig{
		db: dbQueries,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlerHealthz)
	mux.HandleFunc("GET /categories", apiCfg.handlerListCategories)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server is running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
