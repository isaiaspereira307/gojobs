package main

import (
	"github.com/isaiaspereira307/gojobs/config"
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
}
