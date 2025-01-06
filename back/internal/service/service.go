package service

import (
	"database/sql"
	"fmt"
	"log"

	// "github.com/BA999ZAI/QRQuiz/internal/adapter/sqlite"
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
	db, err := initDB(cfg)
	if err != nil {
		log.Println("failed to initialize database: ", err)
	}
	defer db.Close()
	fmt.Println("Successfully connected to the database")

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
	routes.Run(fmt.Sprintf(":%d", cfg.HttpPort))
}

func loadConfig() (*config.Config, error) {
	cfg, err := config.InitConfig()
	if err != nil {
		return nil, fmt.Errorf("error with load config: %v", err)
	}

	return cfg, nil
}

func initDB(cfg *config.Config) (*sql.DB, error) {
	var (
		db      *sql.DB
		err     error
		connStr string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBName,
		)
	)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
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
