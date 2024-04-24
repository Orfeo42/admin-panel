package db

import (
	"fmt"
	"os"

	"github.com/Orfeo42/admin-panel/data"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getConnectionString()), &gorm.Config{})

	if err != nil {
		fmt.Println("error in db connection")
		fmt.Println(err.Error())
		return nil, err
	}

	err = initDB(db)
	if err != nil {
		fmt.Println("error in db update")
		return nil, err
	}

	return db, nil
}

func initDB(db *gorm.DB) error {
	fmt.Println("Updating Schema...")
	err := db.AutoMigrate(
		&data.CustomerModel{},
		&data.InvoiceModel{},
		&data.OrderModel{},
	)
	if err != nil {
		fmt.Println("Error in Updating Schema")
		return err
	}
	fmt.Println("Schema Updated!")
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
