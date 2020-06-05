package tests

import (
	"github.com/leozhao0709/learning/src/tcp/server/logger"
)

func Test() {
	var log = logger.GetLogger()
	log.Debug("debug.....")
}
