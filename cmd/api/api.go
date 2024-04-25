package api

import (
	"CS230832/BeastVehicles/service/parking"
	"CS230832/BeastVehicles/service/vehicle"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{addr: addr, db: db}
}

func (a *ApiServer) Run() error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	vehicleHandler := vehicle.NewHandler(vehicle.NewStore(a.db))
	vehicleHandler.RegisterRoutes(subrouter)
	parkingHandler := parking.NewHandler(parking.NewStore(a.db))
	parkingHandler.RegisterRoutes(subrouter)

	log.Println("Starting server on ", a.addr)
	return http.ListenAndServe(a.addr, router)
}
