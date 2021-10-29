package sql_store

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

func TestDB(t *testing.T, databaseUrl string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			_, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
			if err != nil {
				t.Fatal(err)
			}
		}
		db.Close()
	}
}
