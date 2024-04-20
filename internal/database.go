package internal

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

func NewDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "local.db")

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	err = startDatabase(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func startDatabase(db *sql.DB) error {
	data, err := os.ReadFile("ddl.sql")

	if err != nil {
		return err
	}

	_, err = db.Exec(string(data))

	return err
}
