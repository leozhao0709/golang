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

func init() {
	loggerMode = DevMode
}

// SetLoggerMode set the logger mode
func SetLoggerMode(mode uint8) {
	loggerMode = mode
}

// GetLoggerMode get the logger mode
func GetLoggerMode() uint8 {
	return loggerMode
}

// GetLogger get the logger
func GetLogger() *zap.Logger {
	if logger == nil {
		var err error
		var config zap.Config
		switch loggerMode {
		case DevMode:
			config = zap.NewDevelopmentConfig()
		case ExampleMode:
			logger = zap.NewExample(zap.AddCaller())
			return logger
		default:
			config = zap.NewProductionConfig()
			config.OutputPaths = []string{"log.json"}
		}

		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = config.Build()

		if err != nil {
			log.Fatalln("initial zap logger failed. Exit!")
		}
		return logger
	}
	return logger
}
