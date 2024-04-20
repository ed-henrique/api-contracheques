package internal

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

func NewDatabase(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)

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
