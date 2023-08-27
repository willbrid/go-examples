package main

import (
	"platform/logging"
	"platform/services"
)

func writeMessage(logger logging.Logger) {
	logger.Info("SportsStore")
}

func main() {
	services.RegisterDefaultServices()
	services.Call(writeMessage)
}
