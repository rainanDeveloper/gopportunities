package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error
	db, err = InitializeSqlite()

	if err != nil {
		fmt.Printf("error initializing sqlite: %v", err)
		return err
	}

	return nil
}

func GetSqliteDb() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)

	return logger
}
