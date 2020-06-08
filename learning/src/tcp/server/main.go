package main

import (
	"fmt"
	"io"
	"net"

	"github.com/labstack/gommon/log"
)

func main() {
	log.SetLevel(log.DEBUG)
	log.SetHeader("${time_rfc3339} ${level}")

	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatal("listen error")
	}
	defer listener.Close()
	log.Info("listen start at 8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accep err", err)
		} else {
			go process(conn)
		}
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	log.Infof("Accept one connect %v", conn.RemoteAddr())

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
		log.Infof("client from %v sent message", conn.RemoteAddr())
		fmt.Println(string(buffer[:n]))
	}
}
