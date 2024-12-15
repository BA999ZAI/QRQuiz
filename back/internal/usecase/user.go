package usecase

import (
	"fmt"

	"github.com/BA999ZAI/QRQuiz/internal/entity"
	"github.com/BA999ZAI/QRQuiz/internal/repository/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) GetUserById(id string) (entity.User, error) {
	user, err := u.DB.GetUserById(id)
	if err != nil {
		return entity.User{}, fmt.Errorf("db GetUserById: %w", err)
	}

	response := u.parseUserRepoToBody(user)

	return response, nil
}

func (u *Usecase) GetAllUsers() ([]entity.User, error) {
	users, err := u.DB.GetUserAll()
	if err != nil {
		return nil, fmt.Errorf("db GetUserAll: %w", err)
	}

	response := make([]entity.User, 0)
	for _, user := range users {
		response = append(response, u.parseUserRepoToBody(user))
	}

	return response, nil
}

func (u *Usecase) CreateUser(user entity.User) error {
	user.ID = uuid.New()

	hashPassword, err := u.hashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("hashPassword: %w", err)
	}

	newUser := model.User{
		ID:           user.ID.String(),
		Email:        user.Email,
		HashPassword: hashPassword,
	}
	if err := u.DB.CreateUser(newUser); err != nil {
		return fmt.Errorf("db createUser: %w", err)
	}

	return nil
}

func (u *Usecase) UpdateUser(user entity.User) (entity.User, error) {
	hashPassword, err := u.hashPassword(user.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("hashPassword: %w", err)
	}

	newDataUser := model.User{
		ID:           user.ID.String(),
		Email:        user.Email,
		HashPassword: hashPassword,
	}

	if err := u.DB.UpdateUser(newDataUser); err != nil {
		return entity.User{}, fmt.Errorf("db UpdateUser: %w", err)
	}

	response, err := u.DB.GetUserById(user.ID.String())
	if err != nil {
		return entity.User{}, fmt.Errorf("usecase GetUserById: %w", err)
	}

	return u.parseUserRepoToBody(response), nil
}

func (u *Usecase) DeleteUser(id string) error {
	if err := u.DB.DeleteUser(id); err != nil {
		return fmt.Errorf("db DeleteUser: %w", err)
	}

	return nil
}

func (u *Usecase) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

func (u *Usecase) parseUserRepoToBody(user model.User) entity.User {
	userID, err := uuid.Parse(user.ID)
	if err != nil {
		fmt.Println("uuid.Parse user.ID: ", err)
	}

	response := entity.User{
		ID:       userID,
		Email:    user.Email,
		Password: user.HashPassword,
	}

	return response
}
