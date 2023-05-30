package store

import (
	"github.com/bnallapeta/logo-revelio/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("../data/logo-revelio.db"), &gorm.Config{})
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
