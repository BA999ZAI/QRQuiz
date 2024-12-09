package model

type Quiz struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
}

type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	ID      int    `json:"id"`
	Answer  string `json:"answer"`
	Correct bool   `json:"correct"`
}
