package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BA999ZAI/QRQuiz/internal/service"
)

func main() {
	go service.StartApp()

	waitForShutdown()
}

func waitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down server...")

	os.Exit(0)
}
