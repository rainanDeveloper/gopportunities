package main

import (
	"github.com/rainanDeveloper/gopportunities/config"
	"github.com/rainanDeveloper/gopportunities/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	router.Initialize()
}
