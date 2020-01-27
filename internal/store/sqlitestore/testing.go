package sqlitestore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
	// _ "github.com/mattn/go-sqlite3" // ..
)

//TestDB ..
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("sqlite3", databaseURL)
	if (err != nil) || (db == nil) {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("DELETE FROM %s", strings.Join(tables, ",")))
		}
		db.Close()
	}
}
