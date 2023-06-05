package store

import (
	"os"

	"github.com/bnallapeta/logo-revelio/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "../data/logo-revelio.db" // Fallback to default path if DB_PATH is not set
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitializeDB() (*gorm.DB, error) {
	db, err := ConnectToDB()
	if err != nil {
		return nil, err
	}

	// AutoMigrate the necessary tables
	err = db.AutoMigrate(&model.User{}, &model.Score{}, &model.Answer{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
