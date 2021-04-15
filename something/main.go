package main

import (
	"fmt"
	"regexp"
	"strconv"
)


func main() {

	//sl1 := []int{1, 2, 3, 4, 5, 6}

	phoneRegex, _ := regexp.Compile(`^1(3[0-9]|4[01456879]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\d{8}$`)
	fmt.Println(phoneRegex.MatchString("18296874638"))

}

func iterateWithDel() {
	type d struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var list []d
	for i := 0; i < 10; i++ {
		list = append(list, d{
			Name: "345",
			Age:  i,
		})
	}
	a := 5
	delete := false
	for i := 0; i < len(list); {
		if list[i].Age == a {
			if delete {
				list = append(list[:i], list[i+1:]...)
				continue
			}
			list[i].Name = "haha"
		}
		i++
	}
	fmt.Println(list)
}

func ldkfj() {
	score := int64(341)<<16 | (int64(274877906943 & 0xffff))
	fmt.Println(score, strconv.FormatInt(score, 2))
	fmt.Println(strconv.FormatInt(int64(274877906943&0xffffffff), 2))
	fmt.Println(strconv.FormatInt(int64(341)<<32, 2))
}
