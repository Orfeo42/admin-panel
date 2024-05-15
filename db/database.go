package db

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	databaseInstance *gorm.DB
	lock             = &sync.Mutex{}
)

func GetInstance() (*gorm.DB, error) {
	if isDBCreated() {
		return databaseInstance, nil
	}
	err := createDatabase()
	if err != nil {
		return nil, err
	}
	return databaseInstance, nil
}

func isDBCreated() bool {
	if databaseInstance != nil {
		return true
	}

	lock.Lock()
	defer lock.Unlock()

	return databaseInstance != nil
}

func createDatabase() error {
	newLogger := logger.New(
		log.New("Gorm: "),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)
	connString, err := getConnectionString()
	if err != nil {
		log.Errorf("Error creating connetion string: %+v", err)
		return err
	}
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Errorf("Error in db connection: %+v", err)
		return err
	}
	databaseInstance = db
	return nil
}

func getConnectionString() (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file", err)
		return "", err
	}
	dbHost := getEnv("POSTGRES_HOST", "127.0.0.1")
	dbPort := getEnv("POSTGRES_PORT", "5432")
	dbName := getEnv("POSTGRES_DB", "database")
	dbUser := getEnv("POSTGRES_USER", "user")
	dbPassword := getEnv("POSTGRES_PASSWORD", "password")
	dbSSLMode := getEnv("POSTGRES_SSLMODE", "disable")
	uri := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=%s user=%s password=%s", dbHost, dbPort, dbName, dbSSLMode, dbUser, dbPassword)
	return uri, nil
}

func getEnv(name, defaultValue string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return defaultValue
}
