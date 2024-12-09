package controller

import (
	"net/http"

	"github.com/BA999ZAI/QRQuiz/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Server) handlerQuizPatchById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
	}

	quiz := entity.Quiz{}
	if err := c.BindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parse body: ": err.Error()})
	}

	quiz, err := s.Usecase.UpdateQuiz(quiz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error usecase": err.Error()})
	}

	c.JSON(http.StatusOK, quiz)
}
