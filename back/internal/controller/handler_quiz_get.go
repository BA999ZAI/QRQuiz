package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
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

	codeData, err := generate(quiz.LinkToQuiz, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"generate qr-code": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"quiz": quiz,
		"qr":   codeData,
	})
}

func generate(content string, size int) ([]byte, error) {
	qrCode, err := qrcode.Encode(content, qrcode.Medium, size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	return qrCode, nil
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
