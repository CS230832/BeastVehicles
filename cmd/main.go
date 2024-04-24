package main

import (
	"CS230832/BeastVehicles/cmd/api"
	"CS230832/BeastVehicles/config"
	"CS230832/BeastVehicles/db"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
	})

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err := initStorage(db); err != nil {
		panic(err)
	}

	log.Println("Database connected")

	addr := fmt.Sprintf("%s:%s", config.Envs.Host, config.Envs.Port)

	apiServer := api.NewApiServer(addr, db)

	if err := apiServer.Run(); err != nil {
		panic(err)
	}
}

func initStorage(db *sql.DB) error {
	return db.Ping()
}
