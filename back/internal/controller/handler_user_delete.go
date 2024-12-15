package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handlerUserDeleteById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	if err := s.Usecase.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error usecase": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User is deleted")
}
