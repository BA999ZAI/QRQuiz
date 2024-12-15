package controller

import (
	"net/http"

	"github.com/BA999ZAI/QRQuiz/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) handlerUserPatchById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parse body: ": err.Error()})
		return
	}

	newId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parse id: ": err.Error()})
		return
	}

	user.ID = newId

	user, err = s.Usecase.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error usecase": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
