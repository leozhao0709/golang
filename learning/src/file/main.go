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
	// ioUtilRead()
	// bufWrite()
	copyTest()
}

func copyTest() {
	_, filePath, _, _ := runtime.Caller(1)
	fmt.Println(filePath)
	dirname := filepath.Dir(filePath)
	_, err := copy(filepath.Join(dirname, "./test.txt"), filepath.Join(dirname, "./test1.txt"))

	if err != nil {
		fmt.Println(err)
	}
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
		} else {
			fmt.Print(string(bytes))
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

func bufWrite() {
	_, filePath, _, _ := runtime.Caller(1)
	dirname := filepath.Dir(filePath)
	filePath = filepath.Join(dirname, "./writeAndCreate.txt")

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	str := "Hello World!34345\r\n"

	writer.WriteString(str)
	writer.WriteString(str)

	writer.Flush()
}

func copy(src, dest string) (int64, error) {
	srcFile, err := os.Open(src)

	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	// srcBuffer := bufio.NewReader(srcFile)

	destFile, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return 0, err
	}

	// destBuff := bufio.NewWriter(destFile)
	return io.Copy(destFile, srcFile)
}
