package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	db "github.com/ndbac/go-log/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://golog:golog@localhost:5432/golog?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}
	testQueries = db.New(conn)

	os.Exit(m.Run())
}
