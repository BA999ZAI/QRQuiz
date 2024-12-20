package usecase

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"

	"github.com/BA999ZAI/QRQuiz/internal/entity"
	"github.com/BA999ZAI/QRQuiz/internal/repository/model"
	"github.com/google/uuid"
)

func (u *Usecase) GetQuizById(id string) (entity.Quiz, map[int]map[string]int, error) {
	rawQuiz, err := u.DB.GetQuizById(id)
	if err != nil {
		return entity.Quiz{}, nil, fmt.Errorf("db GetQuizById: %w", err)
	}

	quiz := u.parseQuizRepoToBody(rawQuiz)

	if quiz.Status {
		results := u.calculateResults(quiz)
		return quiz, results, nil
	}

	return quiz, nil, nil
}

func (u *Usecase) calculateResults(quiz entity.Quiz) map[int]map[string]int {
	results := make(map[int]map[string]int) // questionID -> answer -> count

	for _, userReplies := range quiz.Results {
		for _, reply := range userReplies {
			if _, ok := results[reply.ID]; !ok {
				results[reply.ID] = make(map[string]int)
			}
			results[reply.ID][reply.Reply]++
		}
	}

	return results
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
	body.Results = [][]entity.Reply{}
	body.CreatedAt = time.Now()
	body.LinkToQuiz = fmt.Sprintf("http://localhost:3000/quiz/:%s", body.ID.String())

	if body.TimeToLive.IsZero() {
		body.TimeToLive = time.Now().Add(time.Hour * 24 * 7)
	}

	if body.TimeToLive.Before(time.Now()) {
		body.Status = true
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

func (u *Usecase) AddResult(id string, result []entity.Reply) error {
	rawQuiz, err := u.DB.GetQuizById(id)
	if err != nil {
		return fmt.Errorf("db GetQuizById: %w", err)
	}

	quiz := u.parseQuizRepoToBody(rawQuiz)

	newResult := make([]entity.Reply, 0, len(result))
	for _, reply := range result {
		newResult = append(newResult, entity.Reply{
			ID:    reply.ID,
			Reply: reply.Reply,
		})
	}

	quiz.Results = append(quiz.Results, newResult)

	resultsJSON, err := json.Marshal(quiz.Results)
	if err != nil {
		return fmt.Errorf("json.Marshal results: %w", err)
	}

	newQuiz := model.Quiz{
		ID:         rawQuiz.ID,
		Title:      rawQuiz.Title,
		Questions:  rawQuiz.Questions,
		Results:    string(resultsJSON),
		CreatedAt:  rawQuiz.CreatedAt,
		TimeToLive: rawQuiz.TimeToLive,
		LinkToQuiz: rawQuiz.LinkToQuiz,
		Status:     rawQuiz.Status,
		UserID:     rawQuiz.UserID,
	}

	if err := u.DB.AddResultToQuiz(newQuiz); err != nil {
		return fmt.Errorf("db UpdateQuiz: %w", err)
	}

	return nil
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

		if err := u.exportQuizResultsToExcel(val.ID); err != nil {
			fmt.Println("Failed to export quiz results:", err)
		}
	}

	return nil
}

// exportQuizResultsToExcel создает Excel-файл с результатами опроса
func (u *Usecase) exportQuizResultsToExcel(quizID string) error {
	rawQuiz, err := u.DB.GetQuizById(quizID)
	if err != nil {
		return fmt.Errorf("db GetQuizById: %w", err)
	}

	quiz := u.parseQuizRepoToBody(rawQuiz)

	f := excelize.NewFile()
	sheetName := "Results"
	f.NewSheet(sheetName)

	headers := []string{"Answer #", "Answer Value"}
	for col, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+col)
		f.SetCellValue(sheetName, cell, header)
	}

	row := 2
	for _, results := range quiz.Results {
		for _, reply := range results {
			f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), reply.ID)    // Номер ответа
			f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), reply.Reply) // Значение ответа
			row++
		}
	}

	fileName := fmt.Sprintf("quiz_results_%s.xlsx", quizID)
	if err := f.SaveAs(fileName); err != nil {
		return fmt.Errorf("failed to save Excel file: %w", err)
	}

	fmt.Printf("Excel file created: %s\n", fileName)
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
	var results [][]entity.Reply

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
