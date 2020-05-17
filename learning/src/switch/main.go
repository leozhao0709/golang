package main

import "fmt"

func main() {
	var month int8
	fmt.Println("please give a month:")
	_, err := fmt.Scanln(&month)

	if err != nil {
		fmt.Println("Some error occur!", err)
	} else {
		switch month {
		case 3, 4, 5:
			fmt.Println("spring")
		case 6, 7, 8:
			fmt.Println("summer")
		case 9, 10, 11:
			fmt.Println("autumn")
		case 12, 1, 2:
			fmt.Println("winter")
		default:
			fmt.Println("please give a correct month.")
		}
	}
}
