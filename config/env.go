package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string

	DBUser     string
	DBPassword string
	DBName     string
	DBAddress  string

	JWTSecret              string
	JWTExpirationInSeconds int

	AdminEmail string
	AdminPassword string
	AdminFirstName string
	AdminLastName string
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

		JWTSecret:              getEnv("JWT_SECRET", "ahleeafageaGUFGEWrudppifuponaefanpsFGEIUBFIB"),
		JWTExpirationInSeconds: getEnvInt("JWT_EXPIRATION_IN_SECONDS", 60*60*24*30),

		AdminEmail: getEnv("ADMIN_EMAIL", "admin@gmail.com"),
		AdminPassword: getEnv("ADMIN_PASSWORD", "admin"),
		AdminFirstName: getEnv("ADMIN_FIRST_NAME", "admin"),
		AdminLastName: getEnv("ADMIN_LAST_NAME", "admin"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if conv, err := strconv.Atoi(value); err != nil {
			return fallback
		} else {
			return conv
		}
	}

	return fallback
}
