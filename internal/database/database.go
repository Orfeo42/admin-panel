package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

var (
	dbInstance *gorm.DB
	lock       = &sync.Mutex{}
)

func DBInstance() *gorm.DB {
	if isDBCreated() {
		return dbInstance
	}

	createDatabase()

	return dbInstance
}

func isDBCreated() bool {
	if dbInstance != nil {
		return true
	}

	lock.Lock()
	defer lock.Unlock()

	return dbInstance != nil
}

func createDatabase() {
	logLocal := logger.New(
		createNewLogger(),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)
	connString, err := getConnectionString()
	if err != nil {
		log.Fatal("Error creating connection string", err)
	}
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{Logger: logLocal})
	if err != nil {
		log.Fatal("Error in db connection", err)
	}
	dbInstance = db
}

func createNewLogger() *log.Logger {
	file, err := os.OpenFile("gorm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error closing log file", err)
		}
	}(file)

	return log.New(file, "Gorm: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func getConnectionString() (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
		return "", err
	}
	uri := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=%s user=%s password=%s", host, port, database, "disable", username, password)
	return uri, nil
}
