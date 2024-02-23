package config

import (
	"os"

	"github.com/rainanDeveloper/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.sqlite"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Error("Error during database file creation")
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Error("Error during database file creation")
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite oppening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opportunity{})

	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
