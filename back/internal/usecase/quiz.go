package usecase

import (
	"encoding/json"
	"fmt"
	"time"

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

func (u *Usecase) GetQuizByUserId(id string) ([]entity.Quiz, error) {
	quizzes, err := u.DB.GetQuizByUserId(id)
	if err != nil {
		return nil, fmt.Errorf("db GetQuizById: %w", err)
	}

	var response []entity.Quiz
	for _, quiz := range quizzes {
		response = append(response, u.parseQuizRepoToBody(quiz))
	}

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

func (u *Usecase) CreateQuiz(body entity.Quiz) error {
	body.ID = uuid.New()
	if body.Results == nil {
		body.Results = []entity.Reply{}
	}
	body.CreatedAt = time.Now()
	if body.TimeToLive.IsZero() {
		body.TimeToLive = time.Now().Add(time.Hour * 24 * 7)
	}

	newQuiz := u.parseQuizBodyToRepo(body)
	if err := u.DB.CreateQuiz(newQuiz); err != nil {
		return fmt.Errorf("db createQuiz: %w", err)
	}

	return nil
}

func (u *Usecase) DeleteQuiz(id string) error {
	if err := u.DB.DeleteQuiz(id); err != nil {
		return fmt.Errorf("db DeleteQuiz: %w", err)
	}

	return nil
}

func (u *Usecase) AddResult(id string, result entity.Reply) (entity.Quiz, error) {
	rawQuiz, err := u.DB.GetQuizById(id)
	if err != nil {
		return entity.Quiz{}, fmt.Errorf("db GetQuizById: %w", err)
	}

	quiz := u.parseQuizRepoToBody(rawQuiz)

	var newResult = make([]entity.Reply, 0, len(quiz.Results)+1)
	newResult = append(newResult, quiz.Results...)
	newResult = append(newResult, result)

	resultsJSON, err := json.Marshal(newResult)
	if err != nil {
		return entity.Quiz{}, fmt.Errorf("json.Marshal results: %w", err)
	}

	newQuiz := model.Quiz{
		ID:         quiz.ID.String(),
		Title:      quiz.Title,
		Questions:  rawQuiz.Questions,
		Results:    string(resultsJSON),
		CreatedAt:  rawQuiz.CreatedAt,
		TimeToLive: rawQuiz.TimeToLive,
		LinkToQuiz: rawQuiz.LinkToQuiz,
		Status:     rawQuiz.Status,
		UserID:     rawQuiz.UserID,
	}

	if err := u.DB.AddResultToQuiz(newQuiz); err != nil {
		return entity.Quiz{}, fmt.Errorf("db UpdateQuiz: %w", err)
	}

	newQuiz, err = u.DB.GetQuizById(id)
	if err != nil {
		return entity.Quiz{}, fmt.Errorf("usecase GetQuizById: %w", err)
	}

	response := u.parseQuizRepoToBody(newQuiz)

	return response, nil
}

func (u *Usecase) CheckQuiz() error {
	quizzes, err := u.DB.GetQuizByStatus()
	if err != nil {
		return fmt.Errorf("db GetQuizByStatus: %w", err)
	}

	for _, val := range quizzes {
		if val.Status {
			continue
		}

		if !val.TimeToLive.Before(time.Now()) {
			continue
		}

		if err := u.DB.UpdateQuizStatus(val.ID, true); err != nil {
			return fmt.Errorf("db UpdateQuiz: %w", err)
		}
	}

	return nil
}

func (u *Usecase) parseQuizBodyToRepo(quiz entity.Quiz) model.Quiz {
	questionsJSON, err := json.Marshal(quiz.Questions)
	if err != nil {
		fmt.Println("json.Marshal questions: ", err)
	}

	resultsJSON, err := json.Marshal(quiz.Results)
	if err != nil {
		fmt.Println("json.Marshal results: ", err)
	}

	newQuiz := model.Quiz{
		ID:         quiz.ID.String(),
		Title:      quiz.Title,
		Questions:  string(questionsJSON),
		Results:    string(resultsJSON),
		CreatedAt:  quiz.CreatedAt,
		TimeToLive: quiz.TimeToLive,
		LinkToQuiz: quiz.LinkToQuiz,
		Status:     quiz.Status,
		UserID:     quiz.UserID.String(),
	}

	return newQuiz
}

func (u *Usecase) parseQuizRepoToBody(quiz model.Quiz) entity.Quiz {
	var questions []entity.Question
	var results []entity.Reply

	if err := json.Unmarshal([]byte(quiz.Questions), &questions); err != nil {
		fmt.Println("json.Unmarshal questions: ", err)
	}

	if err := json.Unmarshal([]byte(quiz.Results), &results); err != nil {
		fmt.Println("json.Unmarshal results: ", err)
	}

	quizID, err := uuid.Parse(quiz.ID)
	if err != nil {
		fmt.Println("uuid.Parse quiz.ID: ", err)
	}

	userID, err := uuid.Parse(quiz.UserID)
	if err != nil {
		fmt.Println("uuid.Parse quiz.UserID: ", err)
	}

	response := entity.Quiz{
		ID:         quizID,
		Title:      quiz.Title,
		Questions:  questions,
		Results:    results,
		CreatedAt:  quiz.CreatedAt,
		TimeToLive: quiz.TimeToLive,
		LinkToQuiz: quiz.LinkToQuiz,
		Status:     quiz.Status,
		UserID:     userID,
	}

	return response
}
