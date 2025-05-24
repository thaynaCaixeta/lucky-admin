package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	// Server config default values and keys
	defaultServerPort = "8080"
	serverPortKey     = "SERVER_PORT"

	defaultAddr = "localhost"
	addrKey     = "SERVER_ADDR"

	// Postgres config default values and keys
	defaultDbUser = "lucky-admin-api"
	dbUserKey     = "DB_USER"

	defaultDbPass = "123456"
	dbPassKey     = "DB_PASSWORD"

	defaultDbHost = "localhost"
	dbHostKey     = "DB_HOST"

	defaultDbPort = "5432"
	dbPortKey     = "DB_PORT"

	defaultDbName = "lucky-admin"
	dbNameKey     = "DB_NAME"

	defaultDbSSLMode = "disable"
	dbSSLModeKey     = "DB_SSLMODE"
)

type AppConfig struct {
	PostgresConfig PostgresConfig
	ServerConfig   ServerConfig
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SslMode  string
}

type ServerConfig struct {
	Addr string
	Port string
}

func NewAppConfig() AppConfig {
	loadEnvConfig()

	return AppConfig{
		PostgresConfig: PostgresConfig{
			User:     getEnvOrDefaultAsString(defaultDbUser, dbUserKey),
			Password: getEnvOrDefaultAsString(defaultDbPass, dbPassKey),
			Host:     getEnvOrDefaultAsString(defaultDbHost, dbHostKey),
			Port:     getEnvOrDefaultAsString(defaultDbPort, dbPortKey),
			Name:     getEnvOrDefaultAsString(defaultDbName, dbNameKey),
			SslMode:  getEnvOrDefaultAsString(defaultDbSSLMode, dbSSLModeKey),
		},
		ServerConfig: ServerConfig{
			Addr: getEnvOrDefaultAsString(defaultAddr, addrKey),
			Port: getEnvOrDefaultAsString(defaultServerPort, serverPortKey),
		},
	}
}

func loadEnvConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}
}

func getEnvOrDefaultAsString(defValue, key string) string {
	envValue := os.Getenv(key)
	if envValue == "" {
		return defValue
	}
	return envValue
}

/* func getEnvOrDefaultAsInt(defValue, key string) int {
	envValue := os.Getenv(key)
	if envValue == "" {
		return toInt(defValue)
	}
	return toInt(envValue)
}

func toInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		//TODO handle error properly to avoid crash
		log.Fatalf("Error while parsing .env configs: %v", err)
	}
	return intValue
} */
