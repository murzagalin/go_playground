package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/murzagalin/go_playground/api"
	db "github.com/murzagalin/go_playground/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	// Create a new server
	store := db.NewStore(conn)
	server := api.NewServer(store)

	// Start and run the server
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
