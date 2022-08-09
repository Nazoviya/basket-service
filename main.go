package main

import (
	"database/sql"
	"log"

	"github.com/Nazoviya/basketService/api"
	db "github.com/Nazoviya/basketService/db/sqlc"
	"github.com/Nazoviya/basketService/util"
	_ "github.com/lib/pq"
)

func main() {
	// load necessary data into var config.
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}

	// create a connection to database with config properties.
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	// start database connection at given server address.
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
