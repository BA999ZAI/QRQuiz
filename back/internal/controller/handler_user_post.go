package controller

import (
	"log"
	"net/http"

	"github.com/BA999ZAI/QRQuiz/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Server) handlerUserPost(c *gin.Context) {
	body := entity.User{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parse body: ": err.Error()})
		return
	}

	log.Println("body request:", body)

	

	c.JSON(http.StatusCreated, "User is created")
}
