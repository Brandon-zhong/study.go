package main

import (
	"fmt"
	"sort"
)

func main() {
	var l = []User{User{id: 1, name: "fkasd"}, User{id: 2, name: "dfasd"}}
	sort.Sort(list(l))
	testEmptyInterface()
}

func testSort() {
	var l = []User{User{id: 1, name: "fkasd"}, User{id: 2, name: "dfasd"}}
	sort.Sort(list(l))
}

type User struct {
	id   int32
	name string
}

type list []User

func (u list) Len() int {
	return len(u)
}

func (u list) Less(i, j int) bool {
	return u[i].id > u[j].id
}
func (u list) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func testEmptyInterface() {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
