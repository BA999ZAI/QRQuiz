package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/BA999ZAI/QRQuiz/internal/repository/model"
)

func (r *Repository) GetUserById(id string) (model.User, error) {
	query := `SELECT * FROM users WHERE id = ?`
	var user model.User

	if err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.HashPassword,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf("user not found: %w", err)
		}

		return model.User{}, fmt.Errorf("failed to query user by ID: %w", err)
	}

	return user, nil
}

func (r *Repository) GetUserAll() ([]model.User, error) {
	query := `SELECT * FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("db Query: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.HashPassword,
		); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *Repository) CreateUser(user model.User) error {
	if _, err := r.db.Exec(`
		INSERT INTO users (
			id,
			email,
			password
		) VALUES (?, ?, ?)`,
		user.ID,
		user.Email,
		user.HashPassword,
	); err != nil {
		return fmt.Errorf("db Exec: %w", err)
	}

	return nil
}

func (r *Repository) UpdateUser(user model.User) error {
	query := `UPDATE users
			SET
				email = ?,
				password = ?
			WHERE id = ?`
	if _, err := r.db.Exec(
		query,
		user.Email,
		user.HashPassword,
		user.ID,
	); err != nil {
		return fmt.Errorf("db Exec: %w", err)
	}

	return nil
}

func (r *Repository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("db Exec: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	return nil
}

func (r *Repository) GetUserByEmail(email string) (model.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = ?`
	var user model.User

	if err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.HashPassword,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, fmt.Errorf("failed to query user by email: %w", err)
	}

	return user, nil
}
