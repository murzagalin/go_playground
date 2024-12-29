package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/murzagalin/go_playground/api"
	db "github.com/murzagalin/go_playground/db/sqlc"
	"github.com/murzagalin/go_playground/util"
)

func main() {
	// Load config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configs: ", err)
	}

	// Connect to DB
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	// Create a new server
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	// Start and run the server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
