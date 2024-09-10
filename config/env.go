package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	SSLMode    string
}

var Envs = initDBConfig().buildDSN()

func initDBConfig() DBConfig {
	godotenv.Load(".devcontainer/.env")
	return DBConfig{
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
		DBName:     getEnv("DB_NAME", "postgres"),
		SSLMode:    getEnv("SSL_MODE", "disable"),
	}
}

func (c DBConfig) buildDSN() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=%s",
		c.DBUser, c.DBPassword, c.DBAddress, c.DBName, c.SSLMode,
	)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
