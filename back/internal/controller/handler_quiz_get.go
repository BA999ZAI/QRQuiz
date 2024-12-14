package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handlerQuizGetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
	}

	quiz, err := s.Usecase.GetQuizById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error usecase": err.Error()})
	}

	c.JSON(http.StatusOK, quiz)
}

func (s *Server) handlerQuizGetByUserId(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
	}

	quiz, err := s.Usecase.GetQuizByUserId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error usecase": err.Error()})
	}

	c.JSON(http.StatusOK, quiz)
}

func (s *Server) handlerQuizGetAll(c *gin.Context) {
	quizes, err := s.Usecase.GetAllQuizes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error usecase": err.Error()})
	}

	c.JSON(http.StatusOK, quizes)
}
