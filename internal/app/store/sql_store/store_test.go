package sql_store_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "host=localhost dbname=postgres sslmode=disable user=postgres password=pass"
	}

	os.Exit(m.Run())
}
