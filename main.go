package main

import (
	"anonymdevs/database/postgres"
	"anonymdevs/providers"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set in the environment variables")
	}
	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		log.Fatal("POSTGRES_URL is not set in the environment variables")
	}
	postgresPool := postgres.GetConnectionPool(postgresURL)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	jwtProvider := providers.NewJWTProvider()
	txProvider := providers.NewTransactionProvider(postgresPool)
	emailProvider := providers.NewEmailProvider(
		os.Getenv("ELASTIC_EMAIL_KEY"),
		os.Getenv("ELASTIC_EMAIL_SENDER"),
	)
	logger.Info("JWT Provider initialized", "provider", jwtProvider)
	logger.Info("Transaction Provider initialized", "provider", txProvider)
	logger.Info("Email Provider initialized", "provider", emailProvider)
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	server := http.Server{
		Addr:         net.JoinHostPort("0.0.0.0", port),
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	log.Println("Server started on port", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
