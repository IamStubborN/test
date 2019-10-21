package app

import (
	"github.com/IamStubborN/test/config"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/logger/instance"
	"log"
)

func initializeLogger(config *config.Config) logger.Logger {
	l, err := instance.NewZapLogger(config.Logger.Level, config.Logger.EncodingType, config.Logger.OutputPaths)
	if err != nil {
		log.Fatalln(err)
	}

	return l
}
