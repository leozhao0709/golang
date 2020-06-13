package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var selection string

start:
	for {
		fmt.Println("欢迎登录多人聊天系统")
		fmt.Println("1. 登录聊天室")
		fmt.Println("2. 注册用户")
		fmt.Println("3. 退出系统")
		fmt.Println("请选择<1-3>:")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			selection = scanner.Text()
			switch strings.TrimSpace(selection) {
			case "1":
				fmt.Println("登录聊天室")
				break start
			case "2":
				fmt.Println("注册用户")
				break start
			case "3":
				fmt.Println("退出系统")
				return
			default:
				fmt.Println("输入有误, 请重新输入")
			}
		} else {
			fmt.Println("输入有误, 请重新输入")
		}
	}

	switch selection {
	case "1":
		fmt.Println("请输入用户ID")
		scanner := bufio.NewScanner(os.Stdin)
		var userID string
		var password string
		if scanner.Scan() {
			userID = scanner.Text()
		}
		fmt.Println("请输入密码")
		if scanner.Scan() {
			password = scanner.Text()
		}
		login(&userID, &password)
	case "2":
		fmt.Println("...开始注册....")
	default:
		break
	}
}
