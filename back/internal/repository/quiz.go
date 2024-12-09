package repository

import "github.com/BA999ZAI/QRQuiz/internal/repository/model"

func (r *Repository) GetQuizById(id string) (model.Quiz, error) {
	return model.Quiz{}, nil
}

func (r *Repository) GetQuizAll() ([]model.Quiz, error) {
	return []model.Quiz{}, nil
}

func (r *Repository) CreateQuiz(quiz model.Quiz) (model.Quiz, error) {

	return quiz, nil
}

func (r *Repository) UpdateQuiz(quiz model.Quiz) (model.Quiz, error) {
	return model.Quiz{}, nil
}

func (r *Repository) DeleteQuiz(id string) error {
	return nil
}
