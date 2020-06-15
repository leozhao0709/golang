package main

import (
	"bufio"
	"net"
	"os"
	"strings"

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
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			text := scanner.Text()

			if strings.TrimSpace(text) == "exit" {
				break
			}

			n, err := conn.Write([]byte(text))
			if err != nil {
				log.Error("send data error", err)
			}

			log.Infof("client send %d bytes", n)
		}
	}
}
