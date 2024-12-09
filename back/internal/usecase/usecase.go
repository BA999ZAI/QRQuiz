package usecase

import "github.com/BA999ZAI/QRQuiz/internal/repository"

type Usecase struct {
	// Repositories
	DB *repository.Repository
}

func NewUsecase(db *repository.Repository) *Usecase {
	return &Usecase{
		DB: db,
	}
}
