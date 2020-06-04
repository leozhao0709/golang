package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatalln("listen error", err) // fatalln will exit os
	}
	defer listener.Close()
	log.Println("listen start at 8888")

	for {
		log.Println("waiting for client connecting...")
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accep err=", err)
		} else {
			log.Printf("Accept success connect=%v\n", conn)
		}
	}
}
