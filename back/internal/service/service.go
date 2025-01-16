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
	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

func StartApp() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("error with load config: %v", err)
	}

	db, err := initDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err := runMigrations(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	repository := initRepository(db)

	usecase := initUsecase(repository)

	routes, err := initRoutes(cfg, usecase)
	if err != nil {
		log.Fatalf("error with init routes: %v", err)
	}

	log.Printf("Server is running on port %d\n", cfg.HttpPort)
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
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Connected to PostgreSQL successfully")
	return db, nil
}

func runMigrations(db *sql.DB) error {

	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	log.Println("Migrations executed successfully")
	return nil
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
