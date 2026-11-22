package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

// Container contains environment variables for the application, database, cache, token, and http server
type (
	Container struct {
		App  *App
		DB   *DB
		HTTP *HTTP
		Auth *Auth
		AI   *AI
	}
	AI struct {
		BaseURL string
		Timeout time.Duration
	}
	// App contains all the environment variables for the application
	App struct {
		Name string
		Env  string
	}
	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}
	// HTTP contains all the environment variables for the http server
	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}

	Auth struct {
		JWTSecret string
	}
)

// New creates a new container instance
func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	http := &HTTP{
		Env:            os.Getenv("APP_ENV"),
		URL:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	auth := &Auth{
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	baseURL := os.Getenv("AI_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8081"
	}
	seconds := 450
	if value := os.Getenv("AI_TIMEOUT_SECONDS"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed <= 0 || parsed > 3600 {
			return nil, fmt.Errorf("AI_TIMEOUT_SECONDS must be between 1 and 3600")
		}
		seconds = parsed
	}
	return &Container{
		app,
		db,
		http,
		auth,
		&AI{BaseURL: baseURL, Timeout: time.Duration(seconds) * time.Second},
	}, nil
}
