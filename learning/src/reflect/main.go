package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/pkg/errors"
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

type person struct {
	name string
}

type stu struct {
	Name    string
	Age     int
	Teacher *person
}

type stuDto struct {
	Name    string
	Height  int
	Teacher *person
}

func copyProperties(source interface{}, target interface{}) error {
	paramIsValid := func(data interface{}) bool {
		reflectTypeKind := reflect.TypeOf(data).Kind()
		return reflectTypeKind == reflect.Array ||
			reflectTypeKind == reflect.Chan ||
			reflectTypeKind == reflect.Slice ||
			reflectTypeKind == reflect.Map ||
			reflectTypeKind == reflect.Ptr
	}

	if !paramIsValid(source) || !paramIsValid(target) {
		return errors.New("Copy error: source and target must be Array, Chan, Slice, Map or Ptr")
	}

	sourceType := reflect.TypeOf(source).Elem()
	sourceVal := reflect.ValueOf(source).Elem()
	targetType := reflect.TypeOf(target).Elem()
	targetVal := reflect.ValueOf(target).Elem()

	for i := 0; i < sourceType.NumField(); i++ {
		sourceField := sourceType.Field(i)
		targetField, ok := targetType.FieldByName(sourceField.Name)
		if ok {
			if sourceField.Type != targetField.Type {
				log.Printf(`Copy warning: source field "%v" type is %v, but target field "%v" type is %v, so ignored`, sourceField.Name, sourceField.Type, targetField.Name, targetField.Type)
				continue
			}
			targetVal.FieldByName(sourceField.Name).Set(sourceVal.Field(i))
		}
	}

	return nil
}

func testCopy() {
	name := "testName"
	age := 18
	s := stu{Name: name, Age: age, Teacher: &person{name: "teacherName"}}
	sd := stuDto{}

	err := copyProperties(s, &sd)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	fmt.Printf("%+v, %+v\n", s, sd)
	// s.Teacher = &person{name: "newTeacher"}
	s.Teacher.name = "newTeacher"
	fmt.Printf("%+v, %+v\n", s.Teacher.name, sd.Teacher.name)
}

func main() {
	// nilValidTest()
	testCopy()
}
