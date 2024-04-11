package main

import (
	"github.com/isaiaspereira307/gojobs/config"
	"github.com/isaiaspereira307/gojobs/router"
)

var (
	logger *config.Logger
)

func main() {
	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %s", err.Error())
		return
	}

	router.Initialize()
}
