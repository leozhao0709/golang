package main

import (
	"bufio"
	"net"
	"os"

	"github.com/labstack/gommon/log"
)

func main() {
	log.SetLevel(log.DEBUG)
	log.SetHeader("${time_rfc3339} ${level}")

	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatal("client connect error")
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Error("read line error", err)
		}

		if string(line) == "exit" {
			break
		}

		var n int
		n, err = conn.Write(line)
		if err != nil {
			log.Error("send data error", err)
		}

		log.Infof("client send %d bytes", n)
	}
}
