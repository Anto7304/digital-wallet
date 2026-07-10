package server

import (
	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-playground/validator/v10"
	_ "github.com/golang-jwt/jwt/v5"
	_ "github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv"
	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/rs/zerolog"
	_ "golang.org/x/crypto/bcrypt"
)

func main() {
	// Your code here
}
