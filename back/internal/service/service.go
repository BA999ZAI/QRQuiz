package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BA999ZAI/QRQuiz/internal/config"
	"github.com/BA999ZAI/QRQuiz/internal/controller"
	"github.com/BA999ZAI/QRQuiz/internal/repository"
	"github.com/BA999ZAI/QRQuiz/internal/usecase"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	// connect to config
	cfg, err := loadConfig()
	if err != nil {
		log.Println("error with load config: ", err)
	}

	// connect to DB
	err = loadDB()
	if err != nil {
		log.Println("error with load DB: ", err)
	}

	// init repository
	repository := initRepository()

	// init usecases
	usecase := initUsecase(repository)

	// init routes
	routes, err := initRoutes(cfg, usecase)
	if err != nil {
		log.Println("error with init routes: ", err)
	}

	// started service
	routes.Run(fmt.Sprintf(":%s", cfg.HttpServer))
}

func loadConfig() (*config.Config, error) {
	cfg, err := config.InitConfig()
	if err != nil {
		return nil, fmt.Errorf("error with load config: %v", err)
	}

	return cfg, nil
}

// TODO: add DB connection
func loadDB() error {
	return nil
}

func initRepository() *repository.Repository {
	db := &sql.DB{}
	return repository.NewRepository(db)
}

func initUsecase(repository *repository.Repository) *usecase.Usecase {
	return usecase.NewUsecase(repository)
}

func initRoutes(cfg *config.Config, usecase *usecase.Usecase) (*gin.Engine, error) {
	r := gin.Default()
	routes := controller.Server{
		Cfg:     cfg,
		Usecase: usecase,
	}

	routes.RegisterRoutes(r)

	return r, nil
}
