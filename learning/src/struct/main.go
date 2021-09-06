package main

import "fmt"

type person struct {
	name  string
	age   int8
	slice []int8
	map1  map[string]string
	// react *react
	react
}

type pointer struct {
	x int
}

type react struct {
	leftUp, rightDown *pointer
}

func test(x *pointer) (*pointer, error) {
	if x.x == 0 {
		return nil, fmt.Errorf("error")
	}

	return &pointer{x: 10}, nil
}

func main() {
	p1 := person{name: "Lei", age: 28}
	fmt.Printf("%+v\n", p1)

	p1.name = "lei"
	p1.age = 29
	p1.slice = append(p1.slice, 2)
	p1.map1 = map[string]string{}
	p1.map1["123"] = "456"
	fmt.Println(p1)

	pCopy := p1
	pCopy.name = "copy"
	fmt.Println(pCopy.name, p1.name)

	p2 := person{}
	p2.name = "name"
	fmt.Printf("%+v\n", p2)

	p3 := new(person)
	p3.name = "new"
	// fmt.Println(*p3)
	fmt.Printf("%+v\n", *p3)

	p4 := &person{}
	p4.name = "pointer"
	fmt.Println((*p4).name)

	changeName(&p2)
	fmt.Println("changed p2 name:", p2.name)

	r := react{leftUp: &pointer{10}, rightDown: &pointer{30}}
	fmt.Printf("leftUp address: %p, rightDown address: %p \n", r.leftUp, r.rightDown)

	p4.leftUp = &pointer{20}
	fmt.Printf("p4 react leftUp address: %p\n", p4.react.leftUp)
	fmt.Printf("p4 leftUp address: %p\n", p4.leftUp)
	p4.react.leftUp = &pointer{10}
	fmt.Printf("p4 r leftUp address: %p\n", p4.react.leftUp)
	fmt.Printf("p4 leftUp address: %p\n", p4.leftUp)

	r1 := react{}
	v, _ := test(r1.leftUp)
	fmt.Print(v.x)

}

func changeName(p *person) {
	p.name = "changedName"
}
