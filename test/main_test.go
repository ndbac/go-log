package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	sqlc "github.com/ndbac/go-log/src/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://golog:golog@localhost:5432/golog?sslmode=disable"
)

var testQueries *sqlc.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}
	testQueries = sqlc.New(conn)

	os.Exit(m.Run())
}
