package main

import (
	"USI-Service/config"
	"USI-Service/logger"
	"USI-Service/urlmapping"
)

func main() {
	logger := logger.GetLogger()

	// log service starting
	logger.Info("Starting service")

	// config service
	configs := config.New()

	configs.GetEnvConfig()

	urlmapping.Server(configs, logger)
}
