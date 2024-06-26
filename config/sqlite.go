package config

import (
	"os"

	"github.com/isaiaspereira307/gojobs/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing SQLite: %s", err.Error())
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Job{})
	if err != nil {
		logger.Errorf("Error migrating SQLite: %s", err.Error())
		return nil, err
	}
	return db, nil
}
