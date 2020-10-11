package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// useArgs()
	useFlag()
}

func useArgs() {
	args := os.Args
	fmt.Println("total", len(args))
	for index, value := range args {
		fmt.Printf("arg[%d]: %v\n", index, value)
	}
}

func useFlag() {
	var user = flag.String("u", "default user", "user description")
	var pwd = flag.String("pwd", "12345", "password description")
	var host = flag.String("h", "127.0.0.1", "host description")
	var port *int = flag.Int("p", 3306, "port description")

	flag.Parse()
	flag.PrintDefaults()
	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v\n", *user, *pwd, *host, *port)
}
