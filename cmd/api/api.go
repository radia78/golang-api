package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/radia78/personal-finance-api/service/user"
)

// Create an API server struct
type APIServer struct {
	addr string
	db   *sql.DB
}

// Takes the address and database and give it to the API server pointer
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Create a new handler for the user service
	userService := user.NewHandler()
	userService.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
