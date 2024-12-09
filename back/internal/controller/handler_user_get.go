package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handlerUserGetById(c *gin.Context) {
	c.JSON(http.StatusOK, "User")
}

func (s *Server) handlerUserGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, "User")
}
