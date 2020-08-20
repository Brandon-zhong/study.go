package main

import (
	"fmt"
	"reflect"
	"regexp"
)

func main() {

	var t T = T{Age: 20, Nested: Nested{Email: "asd@foxmail.com"}}
	validate(t)

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
	for i := 0; i < typeOf.NumField(); i++ {
		typeField := typeOf.Field(i)
		valueField := valueOf.Field(i)
		tagContent := typeField.Tag.Get("validate")
		fmt.Println("tag --> ", tagContent)
		switch valueField.Kind() {
		case reflect.Int:
			value := valueField.Int()
			fmt.Println(typeField.Name, "  value -->", value)
		case reflect.String:
			value := valueField.String()
			fmt.Println(typeField.Name, "  value --> ", value)
		case reflect.Struct:
		}
	}

	return validateResult, errmsg
}
