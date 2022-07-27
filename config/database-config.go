package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	cloudURL := os.Getenv("CLOUD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", dbHost, dbUser, dbPass, dbName)
	if cloudURL == "true" {
		cloudUser := os.Getenv("DATABASE_USER")
		cloudPass := os.Getenv("DATABASE_PASS")
		cloudHost := os.Getenv("DATABASE_URL")
		cloudName := os.Getenv("DATABASE_NAME")
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", cloudHost, cloudUser, cloudPass, cloudName)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed a connection database")
	}
	err = db.AutoMigrate(&entity.User{}, &entity.ProductCategory{}, &entity.Product{}, &entity.Message{})
	if err != nil {
		return nil
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	err = dbSQL.Close()
	if err != nil {
		return
	}
}
