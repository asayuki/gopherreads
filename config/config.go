package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBFile          string
	LibraryPath     string
	DataPath        string
	HTTPAddr        string
	ScanInterval    int64
	DefaultEmail    string
	DefaultPassword string
	SessionSecret   string
	SessionExp      int64
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DBFile:          getEnv("DB_FILE", "books.db"),
		LibraryPath:     getEnv("LIBRARY_PATH", ""),
		DataPath:        getEnv("DATA_PATH", "data"),
		HTTPAddr:        getEnv("HTTP_ADDR", ":4000"),
		ScanInterval:    getEnvInt("SCAN_INTERVAL", 10),
		DefaultEmail:    getEnv("DEFAULT_EMAIL", ""),
		DefaultPassword: getEnv("DEFAULT_PASSWORD", ""),
		SessionSecret:   getEnv("SESSION_SECRET", "its-a-me-a-secretio"),
		SessionExp:      getEnvInt("SESSION_EXP", 3600),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
