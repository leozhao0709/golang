package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	// bufReadLine()
	ioUtilRead()
}

func bufReadLine() {
	_, filePath, _, _ := runtime.Caller(1)
	dirname := filepath.Dir(filePath)
	file, err := os.Open(filepath.Join(dirname, "./test.txt"))
	if err != nil {
		fmt.Println("file cannot open", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		bytes, prefix, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error", err)
			}
			break
		}
		if !prefix {
			fmt.Println(string(bytes))
		}
	}

	fmt.Println("read finish!!")
}

func ioUtilRead() {
	_, filePath, _, _ := runtime.Caller(1)
	dirname := filepath.Dir(filePath)
	filePath = filepath.Join(dirname, "./test.txt")

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read error", err)
	}
	fmt.Println(string(bytes))
}
