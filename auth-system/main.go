package main

import (
	"auth-system-demo/auth"
	"auth-system-demo/db"
	"auth-system-demo/handlers"
	"auth-system-demo/middleware"
	"auth-system-demo/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	requiredVars := []string{"DATABASE_URL", "JWT_SECRET"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Required env variable %s is not set", v)
		}
	}
}

func main() {
	loadEnv()

	database, err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database")

	r := mux.NewRouter()

	userRepo := models.NewUserRepository(database)
	refreshTokenRepo := models.NewRefreshTokenRepository(database)

	authService := auth.NewAuthService(userRepo, refreshTokenRepo, os.Getenv("JWT_SECRET"), 15*time.Minute)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userRepo)

	// public routes
	r.HandleFunc("/api/auth/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/api/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/auth/refresh", authHandler.RefreshToken).Methods("POST")

	// protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware(authService))
	protected.HandleFunc("/profile", userHandler.Profile).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
