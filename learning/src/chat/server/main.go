package main

import (
	"io"
	"net"

	"github.com/labstack/gommon/log"
)

func main() {
	log.SetLevel(log.DEBUG)
	log.SetHeader("${time_rfc3339} ${level}")

	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatal("server start fail", err)
	}
	defer listener.Close()
	log.Info("listen start at 8888")

	for {
		log.Info("waiting for next connection...")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept err", err)
		}
		log.Info("connect one client with address:", conn.RemoteAddr())
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 4096)
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Error("client read error", err)
			}
			log.Infof("client from %v closed", conn.RemoteAddr())
			return
		}
		log.Infof("client from %v send %v", conn.RemoteAddr(), string(buffer[:n]))
	}
}
