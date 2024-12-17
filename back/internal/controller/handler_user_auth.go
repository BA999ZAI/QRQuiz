package controller

import (
	"github.com/BA999ZAI/QRQuiz/internal/entity"
	"github.com/BA999ZAI/QRQuiz/internal/service/jwt"
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

	token, err := jwt.GenerateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
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

func (s *Server) handlerUserProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to your profile!",
		"userID":  userID,
	})
}

func (s *Server) handlerUserLogout(c *gin.Context) {
	c.JSON(http.StatusOK, "Logout")
}
