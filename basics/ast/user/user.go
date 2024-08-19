package user

import "time"

type User struct {
	name     string
	password string
	age      int
	birthday time.Timer
}

type Teacher struct {
	User
	id *int
}
