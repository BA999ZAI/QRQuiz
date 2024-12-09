package repository

import "github.com/BA999ZAI/QRQuiz/internal/repository/model"

func (r *Repository) GetUserById(id string) (model.User, error) {
	return model.User{}, nil
}

func (r *Repository) GetUserAll() ([]model.User, error) {
	return []model.User{}, nil
}

func (r *Repository) CreateUser(user model.User) (model.User, error) {
	return model.User{}, nil
}

func (r *Repository) UpdateUser(user model.User) (model.User, error) {
	return model.User{}, nil
}

func (r *Repository) DeleteUser(id string) error {
	return nil
}
