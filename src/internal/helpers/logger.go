package helpers

import (
	"aspm.tech/go-boilerplate/src/pkg/logger"
	"github.com/rs/zerolog/log"
)

func ConfigureGlobalLogger() {
	log.Logger = *logger.Configure(logger.Config{
		ConsoleLoggingEnabled: true,
		EncodeLogsAsJson:      true,
		FileLoggingEnabled:    true,
		Directory:             "./tmp",
		Filename:              "go-boilerplate.log",
		MaxSize:               16,
		MaxBackups:            2,
		MaxAge:                1,
	}).Logger
}
