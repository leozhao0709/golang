package main

import (
	"net"

	"github.com/leozhao0709/learning/src/tcp/server/logger"
	"github.com/leozhao0709/learning/src/tcp/server/tests"
	"go.uber.org/zap"
)

var log = logger.GetLogger()

func main() {
	defer log.Sync()
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Error("listen error", zap.Error(err))
	}
	defer listener.Close()
	log.Info("listen start at 8888")

	tests.Test()
	for {
		log.Info("waiting for client connecting...")
		conn, err := listener.Accept()
		if err != nil {
			log.Error("Accep err", zap.Error(err))
		} else {
			log.Info("Accept success connect", zap.Any("conn", conn))
		}
	}
}
