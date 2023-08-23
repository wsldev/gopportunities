package main

import (
	"github.com/wsldev/gopportunities/config"
	"github.com/wsldev/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// Initialize Router
	router.Initialize()
}
