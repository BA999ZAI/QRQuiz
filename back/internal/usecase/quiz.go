package usecase

import (
	"fmt"

	"github.com/BA999ZAI/QRQuiz/internal/entity"
	"github.com/BA999ZAI/QRQuiz/internal/repository/model"
	"github.com/google/uuid"
)

func (u *Usecase) GetQuizById(id string) (entity.Quiz, error) {
	quiz, err := u.DB.GetQuizById(id)
	if err != nil {
		return entity.Quiz{}, fmt.Errorf("db GetQuizById: %w", err)
	}

	response := u.parseQuizRepoToBody(quiz)

	return response, nil
}

func (u *Usecase) GetAllQuizes() ([]entity.Quiz, error) {
	quizes, err := u.DB.GetQuizAll()
	if err != nil {
		return []entity.Quiz{}, fmt.Errorf("db GetQuizAll: %w", err)
	}

	response := make([]entity.Quiz, 0)
	for _, quiz := range quizes {
		response = append(response, u.parseQuizRepoToBody(quiz))
	}

	return response, nil
}

func (u *Usecase) CreateQuiz(body entity.Quiz) (entity.Quiz, error) {
	newQuiz := u.parseQuizBodyToRepo(body)

	quiz, err := u.DB.CreateQuiz(newQuiz)
	if err != nil {
		return entity.Quiz{}, fmt.Errorf("db createQuiz: %w", err)
	}

	response := u.parseQuizRepoToBody(quiz)

	return response, nil
}

func (u *Usecase) UpdateQuiz(quiz entity.Quiz) (entity.Quiz, error) {
	newQuiz := u.parseQuizBodyToRepo(quiz)

	updatedQuiz, err := u.DB.UpdateQuiz(newQuiz)
	if err != nil {
		return entity.Quiz{}, fmt.Errorf("db UpdateQuiz: %w", err)
	}

	response := u.parseQuizRepoToBody(updatedQuiz)

	return response, nil
}

func (u *Usecase) DeleteQuiz(id string) error {
	if err := u.DB.DeleteQuiz(id); err != nil {
		return fmt.Errorf("db DeleteQuiz: %w", err)
	}

	return nil
}

func (u *Usecase) parseQuizBodyToRepo(quiz entity.Quiz) model.Quiz {
	newQuizId := uuid.New()
	newQuiz := model.Quiz{
		ID:          newQuizId.String(),
		Title:       quiz.Title,
		Description: quiz.Description,
		Questions:   make([]model.Question, 0),
	}

	newQuestionId := 0
	for indexQuestion, valueQuestion := range quiz.Questions {
		newQuiz.Questions = append(newQuiz.Questions, model.Question{
			ID:       newQuestionId,
			Question: valueQuestion.Question,
			Answers:  make([]model.Answer, 0),
		})

		newAnswerId := 0
		for _, valAnswer := range valueQuestion.Answers {
			newQuiz.Questions[indexQuestion].Answers = append(newQuiz.Questions[indexQuestion].Answers, model.Answer{
				ID:      newAnswerId,
				Answer:  valAnswer.Answer,
				Correct: valAnswer.Correct,
			})

			newAnswerId++
		}

		newQuestionId++
	}

	return newQuiz
}

func (u *Usecase) parseQuizRepoToBody(quiz model.Quiz) entity.Quiz {
	rawQuiz := entity.Quiz{
		ID:          quiz.ID,
		Title:       quiz.Title,
		Description: quiz.Description,
		Questions:   make([]entity.Question, 0),
	}

	for indexQuestion, valueQuestion := range quiz.Questions {
		rawQuiz.Questions = append(rawQuiz.Questions, entity.Question{
			ID:       valueQuestion.ID,
			Question: valueQuestion.Question,
			Answers:  make([]entity.Answer, 0),
		})

		for _, valAnswer := range valueQuestion.Answers {
			rawQuiz.Questions[indexQuestion].Answers = append(rawQuiz.Questions[indexQuestion].Answers, entity.Answer{
				ID:      valAnswer.ID,
				Answer:  valAnswer.Answer,
				Correct: valAnswer.Correct,
			})
		}
	}

	return rawQuiz
}
