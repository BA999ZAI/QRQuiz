package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handlerUserLogin(c *gin.Context) {
	c.JSON(http.StatusOK, "Login")
}

func (s *Server) handlerUserRegister(c *gin.Context) {
	c.JSON(http.StatusOK, "Register")
}

func (s *Server) handlerUserLogout(c *gin.Context) {
	c.JSON(http.StatusOK, "Logout")
}
