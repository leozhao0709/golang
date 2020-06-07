package main

import (
	"net"

	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/learning/src/tcp/server/tests"
)

func main() {

	log.SetLevel(log.DEBUG)
	log.SetHeader("${time_rfc3339} ${level} ${prefix}")
	log.SetPrefix("prefix")

	// file, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal("open log file Fail!")
	// }
	// wrt := io.MultiWriter(os.Stdout, file)
	// log.SetOutput(wrt)

	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatal("listen error")
	}
	defer listener.Close()
	log.Info("listen start at 8888")

	tests.Test()
	for {
		log.Info("waiting for client connecting...")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accep err")
		} else {
			log.Info("Accept success connect %v", conn)
		}
	}
}
