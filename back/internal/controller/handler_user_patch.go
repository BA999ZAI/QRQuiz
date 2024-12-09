package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handlerUserPatchById(c *gin.Context) {
	c.JSON(http.StatusOK, "User")
}
