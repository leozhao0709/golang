package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// DevMode create new dev logger
	DevMode = 1 << iota
	// ExampleMode create new example logger (for testing)
	ExampleMode
	// ProdMode create new prod logger
	ProdMode
)

var logger *zap.Logger
var loggerMode uint8

// SetLoggerMode set the logger mode
func SetLoggerMode(mode uint8) {
	loggerMode = mode
}

// GetLogger get the global logger
func GetLogger() *zap.Logger {
	if logger == nil {
		initLogger()
	}

	return logger
}

func initLogger() {
	var err error
	var config zap.Config
	switch loggerMode {
	case DevMode:
		config = zap.NewDevelopmentConfig()
	case ExampleMode:
		logger = zap.NewExample(zap.AddCaller())
		return
	case ProdMode:
		config = zap.NewProductionConfig()
		config.OutputPaths = []string{"log.log"}
	default:
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err = config.Build()

	if err != nil {
		log.Panicln("initial zap logger failed. Exit!")
	}
}
