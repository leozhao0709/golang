package tests

import (
	"github.com/leozhao0709/learning/src/tcp/server/logger"
	"go.uber.org/zap"
)

var log *zap.Logger = logger.GetLogger()

// func init() {
// 	logger.SetLoggerMode(logger.DevMode)
// 	log = logger.GetLogger()
// 	fmt.Println("...test init...")
// }

func Test() {
	log.Debug("debug.....")
}
