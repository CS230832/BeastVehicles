package api

import (
	"CS230832/BeastVehicles/service/blocks"
	"CS230832/BeastVehicles/service/parkings"
	"CS230832/BeastVehicles/service/users"
	"CS230832/BeastVehicles/service/vehicles"
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

	userStore := users.NewStore(a.db)

	userHandler := users.NewHandler(userStore)
	parkingHandler := parkings.NewHandler(parkings.NewStore(a.db))
	vehicleHandler := vehicles.NewHandler(vehicles.NewStore(a.db))
	blockHandler := blocks.NewHandler(blocks.NewStore(a.db))

	userHandler.RegisterRoutes(subrouter)
	parkingHandler.RegisterRoutes(subrouter, userStore)
	vehicleHandler.RegisterRoutes(subrouter, userStore)
	blockHandler.RegisterRoutes(subrouter)
	
	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, User-Agent")

			if r.Method == "OPTIONS" {
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	log.Println("Starting server on ", a.addr)
	return http.ListenAndServe(a.addr, corsHandler(router))
}
