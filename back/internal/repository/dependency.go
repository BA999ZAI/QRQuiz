package repository

import "github.com/BA999ZAI/QRQuiz/internal/repository/model"

type IRepository interface {
	GetQuizById(id string) (model.Quiz, error)
	GetQuizByUserId(id string) ([]model.Quiz, error)
	GetQuizAll() ([]model.Quiz, error)
	CreateQuiz(quiz model.Quiz) error
	AddResultToQuiz(quiz model.Quiz) error
	DeleteQuiz(id string) error
	GetQuizByStatus() ([]model.Quiz, error)
	UpdateQuizStatus(id string, status bool) error

	GetUserById(id string) (model.User, error)
	GetUserAll() ([]model.User, error)
	CreateUser(user model.User) error
	UpdateUser(user model.User) error
	DeleteUser(id string) error
	GetUserByEmail(email string) (model.User, error)
}
