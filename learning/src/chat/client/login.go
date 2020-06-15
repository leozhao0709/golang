package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/learning/src/chat/common/message"
)

func login(userID *string, password *string) error {
	fmt.Printf("userId=%v, password=%v\n", *userID, *password)
	// return nil

	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Info("....dial fail", err)
		return err
	}

	var loginMessage = &message.LoginMsg{UserID: userID, Password: password}
	var message = message.Message{Type: message.LoginMsgType, Data: loginMessage}

	result, err := json.Marshal(message)

	if err != nil {
		log.Info("message marshal fail", err)
		return err
	}

	conn.Write(result)

	return nil
}
