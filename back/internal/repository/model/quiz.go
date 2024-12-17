package model

import (
	"time"
)

type Quiz struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Questions  string    `json:"questions"`
	Results    string    `json:"results"`
	CreatedAt  time.Time `json:"created_at"`
	TimeToLive time.Time `json:"time_to_live"`
	LinkToQuiz string    `json:"link_to_quiz"`
	Status     bool      `json:"status"`
	UserID     string    `json:"user_id"`
	Answers    []string  `json:"answers"`
}
