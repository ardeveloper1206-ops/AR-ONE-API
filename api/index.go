package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq" // replace with your DB driver if needed
)

// Package-level resources reused across warm invocations:
var db *sql.DB
var router http.Handler

func init() {
	// Initialize DB once per container (reused across invocations).
	// Set DATABASE_URL in Vercel environment variables.
	dsn := os.Getenv("DATABASE_URL")
	if dsn != "" {
		var err error
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("opening db: %v", err)
		}
		// Tuning for serverless environment:
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(2)
		if err := db.Ping(); err != nil {
			log.Fatalf("ping db: %v", err)
		}
	}

	mux := http.NewServeMux()
	// Register endpoints. Incoming requests will be /api/...
	mux.HandleFunc("/api/hello", helloHandler)
	mux.HandleFunc("/api/users", usersHandler)

	router = mux
}

// Handler is the Vercel entrypoint. Do NOT call ListenAndServe.
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from Vercel Go serverless!\n"))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		http.Error(w, "DB not configured", http.StatusInternalServerError)
		return
	}
	// Example placeholder — replace with your real logic
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("users handler (implement DB logic)\n"))
}
