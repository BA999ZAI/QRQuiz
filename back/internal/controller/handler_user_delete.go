package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handlerUserDeleteById(c *gin.Context) {
	c.JSON(http.StatusOK, "User")
}
