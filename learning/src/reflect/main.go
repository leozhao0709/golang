package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func reflectTypeTest() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	var f int
	var g []int
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32
	reflectType(f) // type:int kind:int
	reflectType(g) // type: kind:slice

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "test",
		age:  18,
	}
	var e = book{title: "test"}
	reflectType(d) // type:person kind:struct
	reflectType(e) // type:book kind:struct
}

func nilValidTest() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("test IsValid:", reflect.ValueOf(0).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	fmt.Println("....b....", reflect.ValueOf(b).FieldByName("abc"))
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}

type stu struct {
	name string
}

type stuDto struct {
	name string
}

func copyProperties(s *stu, sd *stuDto) {
	sv := reflect.ValueOf(s)
	// st := reflect.TypeOf(s)

	sdv := reflect.ValueOf(sd)
	sdv.Elem().FieldByName("name").Set(sv.FieldByName("name"))
}

func testCopy() {
	s := &stu{name: "student"}
	sd := &stuDto{}

	copyProperties(s, sd)

	fmt.Println("...student...", s)
	fmt.Println("...studentDto...", sd)
}

func main() {
	// nilValidTest()
	testCopy()
}
