package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"gocloud.dev/postgres"
	_ "gocloud.dev/postgres/gcppostgres"
)

func ConnectAndExec(connStr string, command string) error {
	db, close, err := openDB(connStr)
	if err != nil {
		return err
	}
	defer close()

	return execCommand(db, command)
}

func openDB(connstring string) (db *sql.DB, close func() error, err error) {
	db, err = postgres.Open(context.Background(), connstring)
	if err != nil {
		return nil, nil, fmt.Errorf("connecting to database: %w: %s", err, expectedConnStr)
	}
	if err := db.Ping(); err != nil {
		return nil, nil, fmt.Errorf("pinging database: %w: %s", err, expectedConnStr)
	}
	return db, db.Close, nil
}

func execCommand(db *sql.DB, command string) error {
	_, err := db.Exec(command)
	if err != nil {
		var pqError *pq.Error
		if errors.As(err, &pqError) {
			return fmt.Errorf("execing sql query: %v, %v", pqError.Code, pqError.Message)
		}
		return fmt.Errorf("other exec query: %w", err)
	}
	return nil
}
