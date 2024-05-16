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

	JWTSecret string

	RootUserName  string
	RootPassword  string
	RootFirstName string
	RootLastName  string
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

		JWTSecret: getEnv("JWT_SECRET", "ahleeafageaGUFGEWrudppifuponaefanpsFGEIUBFIB"),

		RootUserName:  getEnv("ROOT_USERNAME", "root"),
		RootPassword:  getEnv("ROOT_PASSWORD", "R0Otp4ssw0rd"),
		RootFirstName: getEnv("ROOT_FIRST_NAME", "Beast"),
		RootLastName:  getEnv("ROOT_LAST_NAME", "Vehicles"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
