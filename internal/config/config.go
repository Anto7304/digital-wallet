package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	Redis     RedisConfig
	App       AppConfig
	JWT       JWTConfig
	RateLimit RateLimitConfig
}

type ServerConfig struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DatabaseConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBName       string
	SSLMode      string
	MaxConns     int
	MinConns     int
	MaxIdleTime  time.Duration
	HealthPeriod time.Duration
}

type JWTConfig struct {
	Secret   string
	Expiry   time.Duration
	Issuer   string
	Audience string
}

type AppConfig struct {
	Environment string
	Debug       bool
	LogLevel    string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type RateLimitConfig struct {
	Request int
	Window  time.Duration
}

// load configuration
func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := &Config{
		Server: ServerConfig{
			Port:         getEnvAsInt("SERVER_PORT", 8080),
			ReadTimeout:  getEnvAsDuration("SERVER_READ_TIMEOUT", "15s"),
			WriteTimeout: getEnvAsDuration("SERVER_WRITE_TIMEOUT", "15s"),
			IdleTimeout:  getEnvAsDuration("SERVER_IDLE_TIMEOUT", "60s"),
		},

		Database: DatabaseConfig{
			Host:         getEnv("DATABASE_HOST", "localhost"),
			Port:         getEnvAsInt("DATABASE_PORT", 5432),
			User:         getEnv("DATABASE_USER", "walletuser"),
			Password:     getEnv("DATABASE_PASSWORD", ""),
			DBName:       getEnv("DATABASE_NAME", "walletdb"),
			SSLMode:      getEnv("DB_SSL_MODE", "disable"),
			MaxConns:     getEnvAsInt("DB_MAX_CONN", 25),
			MinConns:     getEnvAsInt("DB_MIN_CONN", 5),
			MaxIdleTime:  getEnvAsDuration("DB_MAX_IDLE_TIME", "15m"),
			HealthPeriod: getEnvAsDuration("DB_HEALTH_PERIOD", "1m"),
		},

		JWT: JWTConfig{
			Secret:   getEnv("JWT_SECRET", ""),
			Expiry:   getEnvAsDuration("JWT_EXPIRY", "24h"),
			Issuer:   getEnv("JWT_ISSUER", "walle-system"),
			Audience: getEnv("JWT_AUDIENCE", "wallet-user"),
		},

		App: AppConfig{
			Environment: getEnv("APP_ENV", "development"),
			Debug:       getEnvBool("APP_DEBUG", "true"),
			LogLevel:    getEnv("APP_LOG_LEVEL", "debug"),
		},

		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "development"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},

		RateLimit: RateLimitConfig{
			Request: getEnvAsInt("RATE_LIMIT_REQUEST", 100),
			Window:  getEnvAsDuration("RATE_LIMIT_WINDOW", "1m"),
		},
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}
func (c *Config) validate() error {
	if c.JWT.Secret == "" {
		return fmt.Errorf("JWT_SECRET required")
	}

	if len(c.JWT.Secret) < 32 {
		return fmt.Errorf("JWT_SECRET must be at least 32 character")
	}

	if c.App.Environment == "production" && c.Database.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required in production environment")
	}
	return nil
}

// helper function
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue string) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	duration, _ := time.ParseDuration(defaultValue)
	return duration
}

func getEnv(key string, defaultValue string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue string) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	defVal, _ := strconv.ParseBool(defaultValue)
	return defVal
}
