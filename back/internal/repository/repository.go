package repository

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Close() error {
	if err := r.db.Close(); err != nil {
		return err
	}

	return nil
}
