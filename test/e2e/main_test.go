package e2e

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

var (
	db     *sql.DB
	dbPath = "../../app.db"
)

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("failed to open db: %v", err.Error())
	}

	_, err = db.Exec(`DELETE FROM users`)
	if err != nil {
		log.Fatalf("failed to clear user table: %v", err.Error())
	}
	_, err = db.Exec(`UPDATE SQLITE_SEQUENCE SET SEQ=0 WHERE NAME='users'`)
	if err != nil {
		log.Fatalf("failed to reset auto increment on users table: %v", err.Error())
	}

	_, err = db.Exec(`DELETE FROM roles`)
	if err != nil {
		log.Fatalf("failed to clear roles table: %v", err.Error())
	}
	_, err = db.Exec(`UPDATE SQLITE_SEQUENCE SET SEQ=0 WHERE NAME='roles'`)
	if err != nil {
		log.Fatalf("failed to reset auto increment on roles table: %v", err.Error())
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}
