package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	//Initialize db
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("Error initializing SQLite: %s", err.Error())
	}
	return nil
}

func GetSqile() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
