package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handlerQuizDeleteById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
	}

	if err := s.Usecase.DeleteQuiz(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error usecase": err.Error()})
	}

	c.JSON(http.StatusOK, "Quiz is deleted")
}
