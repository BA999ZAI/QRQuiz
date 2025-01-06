package controller

import (
	"github.com/BA999ZAI/QRQuiz/internal/service/middleware"
	// "log"
	"time"

	"github.com/BA999ZAI/QRQuiz/internal/config"
	"github.com/BA999ZAI/QRQuiz/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Cfg     *config.Config
	Usecase *usecase.Usecase
}

func (s *Server) RegisterRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://example.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	group := r.Group(s.Cfg.HttpPrefix)
	{
		// Quizes
		group.GET("/quiz", s.handlerQuizGetAll)      // is working
		group.GET("/quiz/:id", s.handlerQuizGetById) // is working
		group.GET("/quiz/user/:id", s.handlerQuizGetByUserId)
		group.POST("/quiz", s.handlerQuizPost)               // is working
		group.PATCH("/quiz/:id", s.handlerQuizAddResultPost) // is working
		group.DELETE("/quiz/:id", s.handlerQuizDeleteById)   // is working

		// Users
		group.GET("/user", s.handlerUserGetAll)            // is working
		group.GET("/user/:id", s.handlerUserGetById)       // is working
		group.POST("/user", s.handlerUserPost)             // is working
		group.PATCH("/user/:id", s.handlerUserPatchById)   // is working
		group.DELETE("/user/:id", s.handlerUserDeleteById) // is working

		// Authorization
		group.POST("/login", s.handlerUserLogin)
		group.POST("/register", s.handlerUserRegister)

		// Middleware
		auth := group
		auth.Use(middleware.JWTAuthMiddleware())
		// TODO: продумаьть и обернуть пользовательские ручки  в middleware
		//{
		//	auth.GET("/quiz/user/:id", s.handlerQuizGetByUserId) // Пример защищенного маршрута
		//	auth.POST("/quiz", s.handlerQuizPost)                // Создание квиза
		//	auth.PATCH("/quiz/:id", s.handlerQuizAddResultPost)  // Добавление результата
		//	auth.DELETE("/quiz/:id", s.handlerQuizDeleteById)    // Удаление квиза
		//}

		// Health
		group.GET("/health", s.handlerHealth) // is working
	}

	// go func() {
	// 	ticker := time.NewTicker(1 * time.Second)
	// 	defer ticker.Stop()

	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			if err := s.Usecase.CheckQuiz(); err != nil {
	// 				log.Println("error with check quiz: ", err)
	// 			}
	// 		}
	// 	}
	// }()
}
