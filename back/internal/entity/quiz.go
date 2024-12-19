package entity

import (
	"time"

	"github.com/google/uuid"
)

type Quiz struct {
	ID         uuid.UUID  `json:"id"`
	Title      string     `json:"title"`
	Questions  []Question `json:"questions"`
	Results    [][]Reply  `json:"results"`
	CreatedAt  time.Time  `json:"created_at"`
	TimeToLive time.Time  `json:"time_to_live"`
	LinkToQuiz string     `json:"link_to_quiz"`
	Status     bool       `json:"status"`
	UserID     uuid.UUID  `json:"user_id"`
}

type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

type Reply struct {
	ID    int    `json:"id"`
	Reply string `json:"reply"`
}
