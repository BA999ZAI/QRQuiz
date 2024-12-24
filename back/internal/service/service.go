package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BA999ZAI/QRQuiz/internal/adapter/sqlite"
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
	db, err := sql.Open("sqlite3", cfg.DBPATH)
	if err != nil {
		log.Println("failed to initialize database: ", err)
	}
	defer db.Close()

	// init migrations
	if err := sqlite.RunMigrations(db); err != nil {
		log.Println("failed to run migrations: ", err)
	}

	// init repository
	repository := initRepository(db)

	// init usecases
	usecase := initUsecase(repository)

	// init routes
	routes, err := initRoutes(cfg, usecase)
	if err != nil {
		log.Println("error with init routes: ", err)
	}

	// started service
	routes.Run(":10000")
}

func loadConfig() (*config.Config, error) {
	cfg, err := config.InitConfig()
	if err != nil {
		return nil, fmt.Errorf("error with load config: %v", err)
	}

	return cfg, nil
}

func initRepository(db *sql.DB) *repository.Repository {
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
