package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/ndbac/go-log/src/api"
	db "github.com/ndbac/go-log/src/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://golog:golog@localhost:5432/golog?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
