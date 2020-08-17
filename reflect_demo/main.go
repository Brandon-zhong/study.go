package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var not NotknownType = NotknownType{s1: "hahah", S2: "heheh", s3: "foiqur", Age: 20}
	fmt.Printf("main address --> %p\n", &not)
	fmt.Println(not)
	value := reflect.ValueOf(&not).Elem()
	var typeOf = value.Type()
	fmt.Printf("values --> %v, type --> %v\n", value, typeOf)
	for i := 0; i < value.NumField(); i++ {
		var f = value.Field(i)
		fmt.Printf("value --> %v, type --> %v, typeOf --> %v \n", f, f.Type().Name(), typeOf.Field(i).Name)
	}
	value.FieldByName("Age").SetInt(32)
	value.FieldByName("S2").SetString("shengming.zhong")
	fmt.Println(value)
}

type NotknownType struct {
	s1, S2, s3 string
	Age        int32
}

func (n NotknownType) String() string {
	//fmt.Printf("string address --> %p\n", &n)
	return n.s1 + " - " + n.S2 + " - " + n.s3 + " - " + strconv.Itoa(int(n.Age))
}
