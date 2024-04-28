package db

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(getConnectionString()), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in db connection")
		fmt.Println(err.Error())
		return err
	}
	databaseInstance = db
	return nil
}

func getConnectionString() string {
	godotenv.Load()
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "database")
	dbUser := getEnv("DB_USER", "user")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")
	uri := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=%s user=%s password=%s", dbHost, dbPort, dbName, dbSSLMode, dbUser, dbPassword)
	return uri
}

func getEnv(name, defaultValue string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return defaultValue
}
