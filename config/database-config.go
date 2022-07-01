package config

import (
	"fmt"
	"github.com/my-way-teams/my_way_backend/entity"
	"os"

	"github.com/joho/godotenv"
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
	herokuUrl := os.Getenv("DATABASE_URL")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", dbHost, dbUser, dbPass, dbName)
	if len(herokuUrl) != 0 {
		dsn = fmt.Sprintf(os.Getenv("DATABASE_URL"))
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed a connection database")
	}
	err = db.AutoMigrate(&entity.Admin{}, &entity.User{},&entity.Country{}, &entity.City{}, &entity.Place{}, &entity.PlaceCategory{}, &entity.Route{})
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