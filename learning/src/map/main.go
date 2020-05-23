package main

import "fmt"

func main() {
	var cities = make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	delete(cities, "no2")
	fmt.Println(cities)

	var heros = map[string]string{}
	heros["hero1"] = "宋江"

	hero1, exist := heros["hero3"]
	fmt.Println(hero1, exist)

	if hero2, exist := heros["hero2"]; exist {
		fmt.Println(hero2)
	}
}
