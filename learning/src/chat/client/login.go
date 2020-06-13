package main

import "fmt"

func login(userID *string, password *string) error {
	fmt.Printf("userId=%v, password=%v\n", *userID, *password)
	return nil
}
