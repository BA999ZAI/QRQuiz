package controller

import (
	"github.com/BA999ZAI/QRQuiz/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handlerUserLogin(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	user, err := s.Usecase.AuthenticateUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}

func (s *Server) handlerUserRegister(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	if err := s.Usecase.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (s *Server) handlerUserLogout(c *gin.Context) {
	c.JSON(http.StatusOK, "Logout")
}
