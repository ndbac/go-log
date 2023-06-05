package testSqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ndbac/go-log/src/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://golog:golog@localhost:5432/golog?sslmode=disable"
)

var testQueries *sqlc.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}
	testQueries = sqlc.New(testDB)

	os.Exit(m.Run())
}
