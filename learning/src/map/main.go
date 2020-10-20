package main

import "fmt"

func mergeMap(map1 map[string]string, map2 map[string]string) map[string]string {
	var resultMap = map[string]string{}
	for key, value := range map1 {
		resultMap[key] = value
	}

	for key, value := range map2 {
		resultMap[key] = value
	}

	return resultMap
}

func main() {
	// var map1 = map[string]string{
	// 	"key1": "val1",
	// 	"key2": "val2",
	// }

	// var map2 = map[string]string{
	// 	"key3": "val3",
	// 	"key4": "val4",
	// }

	// var map3 = mergeMap(map1, map2)
	// fmt.Println(map3)
	var map1 = map[string]string{}
	var key = "key1"
	map1[key] = "val1"
	fmt.Println(map1)

	var map2 map[string][]string

	for key, value := range map2 {
		fmt.Println(key, value)
	}

	var map3 *map[string]string = &map[string]string{}
	(*map3)["a"] = "a"
	fmt.Println(map3)
	map3 = nil
	fmt.Println(map3)
}

func basicMap() {
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
