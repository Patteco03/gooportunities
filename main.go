package main

import (
	"exemple/gooportunities/config"
	"exemple/gooportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config.Init() failed: %v ", err)
		return
	}

	// Initialize the router
	router.Initialize()
}
