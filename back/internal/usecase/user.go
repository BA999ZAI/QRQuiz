package usecase

import (
	"github.com/BA999ZAI/QRQuiz/internal/entity"
)

func (u *Usecase) GetUserById(id string) (entity.User, error) {
	return entity.User{}, nil
}

func (u *Usecase) GetAllUsers() ([]entity.User, error) {
	return []entity.User{}, nil
}

func (u *Usecase) CreateUser(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func (u *Usecase) UpdateUser(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func (u *Usecase) DeleteUser(id string) error {
	return nil
}
