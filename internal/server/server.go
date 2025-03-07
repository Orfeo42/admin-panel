package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	//_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"

	"admin-panel/internal/database"
)

type Server struct {
	port int

	db *gorm.DB
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   database.DBInstance(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
