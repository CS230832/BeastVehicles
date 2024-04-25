package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Config struct {
	Address  string
	Database string
	User     string
	Password string
}

func NewPostgresStorage(cfg Config) (*sql.DB, error) {
	dsn := cfg.formatDSN()
	log.Printf("Database DSN: '%s'\n", dsn)
	return sql.Open("postgres", dsn)
}

func (cfg Config) formatDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Address,
		cfg.Database,
	)
}
