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
		databaseUrl = "host=localhost dbname=rest_api_test sslmode=disable user=postgres password=pass"
	}

	os.Exit(m.Run())
}
