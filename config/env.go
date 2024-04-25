package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string

	DBUser     string
	DBPassword string
	DBName     string
	DBAddress  string
}

var Envs Config = initConfig()

func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return Config{
		Host: getEnv("HOST", "localhost"),
		Port: getEnv("PORT", "8080"),

		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBName:     getEnv("DB_NAME", "beast-vehicles"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
