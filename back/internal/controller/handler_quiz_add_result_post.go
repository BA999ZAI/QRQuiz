package controller

import (
	"net/http"

	"github.com/BA999ZAI/QRQuiz/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Server) handlerQuizAddResultPost(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	var result []entity.Reply
	if err := c.BindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parse body: ": err.Error()})
		return
	}

	if err := s.Usecase.AddResult(id, result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error usecase": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"quiz": "updated"})
}
