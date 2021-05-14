package user

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SayHello() {
	fmt.Println("User", u.Name, u.Age)
}
