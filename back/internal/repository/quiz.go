package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/BA999ZAI/QRQuiz/internal/repository/model"
)

func (r *Repository) GetQuizById(id string) (model.Quiz, error) {
	query := `SELECT * FROM quizzes WHERE id = ?`
	var quiz model.Quiz

	if err := r.db.QueryRow(query, id).Scan(
		&quiz.ID,
		&quiz.Title,
		&quiz.Questions,
		&quiz.Results,
		&quiz.CreatedAt,
		&quiz.TimeToLive,
		&quiz.LinkToQuiz,
		&quiz.Status,
		&quiz.UserID,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Quiz{}, fmt.Errorf("quiz not found: %w", err)
		}

		return model.Quiz{}, fmt.Errorf("failed to query quiz by ID: %w", err)
	}

	return quiz, nil
}

func (r *Repository) GetQuizByUserId(id string) ([]model.Quiz, error) {
	query := `SELECT * FROM quizzes WHERE user_id = ?`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("db Query: %w", err)
	}
	defer rows.Close()

	var quizzes []model.Quiz
	for rows.Next() {
		var quiz model.Quiz
		if err := rows.Scan(
			&quiz.ID,
			&quiz.Title,
			&quiz.Questions,
			&quiz.Results,
			&quiz.CreatedAt,
			&quiz.TimeToLive,
			&quiz.LinkToQuiz,
			&quiz.Status,
			&quiz.UserID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan quiz: %w", err)
		}

		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}

func (r *Repository) GetQuizAll() ([]model.Quiz, error) {
	query := `SELECT * FROM quizzes`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("db Query: %w", err)
	}
	defer rows.Close()

	var quizzes []model.Quiz
	for rows.Next() {
		var quiz model.Quiz
		if err := rows.Scan(
			&quiz.ID,
			&quiz.Title,
			&quiz.Questions,
			&quiz.Results,
			&quiz.CreatedAt,
			&quiz.TimeToLive,
			&quiz.LinkToQuiz,
			&quiz.Status,
			&quiz.UserID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan quiz: %w", err)
		}

		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}

func (r *Repository) CreateQuiz(quiz model.Quiz) error {
	if _, err := r.db.Exec(`
		INSERT INTO quizzes (
			id,
			title,
			questions,
			results,
			created_at,
			time_to_live,
			link_to_quiz,
			status,
			user_id
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			quiz.ID,
			quiz.Title,
			quiz.Questions,
			quiz.Results,
			quiz.CreatedAt,
			quiz.TimeToLive,
			quiz.LinkToQuiz,
			quiz.Status,
			quiz.UserID,
		); err != nil {
		return fmt.Errorf("db Exec: %w", err)
	}

	return nil
}

func (r *Repository) AddResultToQuiz(quiz model.Quiz) error {
	query := `UPDATE quizzes
		SET
			title = ?,
			questions = ?,
			results = ?,
			created_at = ?,
			time_to_live = ?,
			link_to_quiz = ?,
			status = ?,
			user_id = ?
		WHERE id = ?`
	if _, err := r.db.Exec(
		query,
		quiz.Title,
		quiz.Questions,
		quiz.Results,
		quiz.CreatedAt,
		quiz.TimeToLive,
		quiz.LinkToQuiz,
		quiz.Status,
		quiz.UserID,
		quiz.ID,
	); err != nil {
		return fmt.Errorf("db Exec: %w", err)
	}

	return nil
}

func (r *Repository) DeleteQuiz(id string) error {
	query := `DELETE FROM quizzes WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("db Exec: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("quiz with ID %s not found", id)
	}

	return nil
}

func (r *Repository) GetQuizByStatus() ([]model.Quiz, error) {
	query := `SELECT * FROM quizzes WHERE status = ?`

	rows, err := r.db.Query(query, false)
	if err != nil {
		return nil, fmt.Errorf("db Query: %w", err)
	}
	defer rows.Close()

	var quizzes []model.Quiz
	for rows.Next() {
		var quiz model.Quiz
		if err := rows.Scan(
			&quiz.ID,
			&quiz.Title,
			&quiz.Questions,
			&quiz.Results,
			&quiz.CreatedAt,
			&quiz.TimeToLive,
			&quiz.LinkToQuiz,
			&quiz.Status,
			&quiz.UserID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan quiz: %w", err)
		}

		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}

func (r *Repository) UpdateQuizStatus(id string, status bool) error {
	query := `UPDATE quizzes SET status = ? WHERE id = ?`
	if _, err := r.db.Exec(query, status, id); err != nil {
		return fmt.Errorf("db Exec: %w", err)
	}

	return nil
}