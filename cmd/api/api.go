package api

import (
	"CS230832/BeastVehicles/service/users"
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

	userHandler := users.NewHandler(users.NewStore(a.db))

	userHandler.RegisterRoutes(subrouter)
	
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
