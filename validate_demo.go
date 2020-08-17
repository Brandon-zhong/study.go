package main

import (
	"fmt"
	"reflect"
	"regexp"
)

func main() {

	validate(T{Age: 20, Nested: Nested{Email: "asdf@foxmail.com"}})

}

type Nested struct {
	Email string `validate:"email"`
}
type T struct {
	Age    int `validate:"eq=10"`
	Nested Nested
}

func validateEmail(input string) bool {
	if pass, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w+).([a-z]{2,4})$`, input); pass {
		return true
	}
	return false
}

func validate(v interface{}) (bool, string) {

	validateResult := true
	errmsg := "success"
	typeOf := reflect.TypeOf(v)
	valueOf := reflect.ValueOf(v)

	fmt.Printf("typeOf NumField --> %d, valueOf NumField --> %d\n", typeOf.NumField(), valueOf.NumField())

	/*for i := 0; i < typeOf.NumField(); i++ {

	}*/
	return validateResult, errmsg
}
