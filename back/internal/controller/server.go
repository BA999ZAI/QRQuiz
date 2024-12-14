package controller

import (
	"github.com/BA999ZAI/QRQuiz/internal/config"
	"github.com/BA999ZAI/QRQuiz/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Cfg     *config.Config
	Usecase *usecase.Usecase
}

func (s *Server) RegisterRoutes(r *gin.Engine) {
	group := r.Group(s.Cfg.HttpPrefix)
	{
		// Quizes
		group.GET("/quiz", s.handlerQuizGetAll)// is working
		group.GET("/quiz/:id", s.handlerQuizGetById)// is working
		group.GET("/quiz/user/:id", s.handlerQuizGetByUserId)
		group.POST("/quiz", s.handlerQuizPost)// is working
		group.PATCH("/quiz/:id", s.handlerQuizAddResultPost)// is working
		group.DELETE("/quiz/:id", s.handlerQuizDeleteById)// is working

		// Users
		group.GET("/user", s.handlerUserGetAll)
		group.GET("/user/:id", s.handlerUserGetById)
		group.POST("/user", s.handlerUserPost)
		group.PATCH("/user/:id", s.handlerUserPatchById)
		group.DELETE("/user/:id", s.handlerUserDeleteById)

		// Authorization
		group.POST("/login", s.handlerUserLogin)
		group.POST("/register", s.handlerUserRegister)
		group.GET("/logout", s.handlerUserLogout)

		// Health
		group.GET("/health", s.handlerHealth)// is working
	}
}
