package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handlerHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "The server is alive.")
}
