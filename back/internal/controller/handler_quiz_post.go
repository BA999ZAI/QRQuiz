package controller

import (
	"net/http"

	"github.com/BA999ZAI/QRQuiz/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Server) handlerQuizPost(c *gin.Context) {
	var quiz entity.Quiz
	if err := c.BindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parse body: ": err.Error()})
		return
	}

	if err := s.Usecase.CreateQuiz(quiz); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error usecase": err.Error()})
	}

	c.JSON(http.StatusCreated, "quiz is created")
}
