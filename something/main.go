package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"study.go/util"
)

func main() {

	type A struct {
		Name string `json:"name"`
	}

	type B struct {
		*A
		Age int `json:"age"`
	}

	var d = B{Age: 12, A: &A{Name: "zhong"}}
	bytes := util.MustByte(json.Marshal(d))

	fmt.Println(string(bytes))
	var a A
	json.Unmarshal(bytes, &a)
	fmt.Println(a)

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
