package main

import (
	"database/sql"
	"log"

	"github.com/radia78/personal-finance-api/cmd/api"
	"github.com/radia78/personal-finance-api/config"
	"github.com/radia78/personal-finance-api/db"
)

func main() {

	// Initiate the database
	db, err := db.NewPostgreSQLStorage(config.Envs)
	if err != nil {
		log.Fatal(err)
	}
	// Connect to the database
	initStorage(db)

	// Initiate the server with the pre-specified parameters (we want to move this probably to an env file tbh)
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
