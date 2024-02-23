package main

import (
	"github.com/rainanDeveloper/gopportunities/config"
	"github.com/rainanDeveloper/gopportunities/router"
)

func main() {
	error := config.Init()
	if error != nil {
		panic(error)
	}
	router.Initialize()
}
