package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	tableUsers = `
		CREATE TABLE users (
		id TEXT PRIMARY KEY NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	tableFolders = `CREATE TABLE folders (
		id TEXT PRIMARY KEY NOT NULL UNIQUE,
		name TEXT NOT NULL,
		user_id TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id),
	);`
	tableQuizzes = `CREATE TABLE quizzes (
		id TEXT PRIMARY KEY NOT NULL UNIQUE,
		title TEXT NOT NULL,
		questions TEXT NOT NULL,
		results TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		time_to_live TIMESTAMP DEFAULT (datetime ('now', '+1 day')),
		link_to_quiz TEXT NOT NULL,
		status BOOLEAN NOT NULL DEFAULT FALSE,
		folder_id TEXT NOT NULL,
		FOREIGN KEY (folder_id) REFERENCES folders(id),
		user_id TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id),
	);`
)

func RunMigrations(db *sql.DB) error {
	existsUsers, err := tableExist(db, "users")
	if err != nil {
		return fmt.Errorf("existTable: %w", err)
	}

	if !existsUsers {
		db.Exec(tableUsers)
	}

	existsQuizzes, err := tableExist(db, "quizzes")
	if err != nil {
		return fmt.Errorf("existTable: %w", err)
	}

	if !existsQuizzes {
		db.Exec(tableQuizzes)
	}

	return nil
}

func tableExist(db *sql.DB, table string) (bool, error) {
	query := `SELECT name FROM sqlite_master WHERE type='table' AND name=?`

	var name string
	if err := db.QueryRow(query, table).Scan(&name); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
