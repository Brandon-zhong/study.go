package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {

	var wait sync.WaitGroup

}
func do(l *sync.Mutex) {
}

func getAccountIdFromUserIdAndIndex(userId, index int) int {
	return userId | (index << 27 & 0xffffffff)
}
func getUserIdAndIndexFromAccountId(accountId int) (userId, index int) {
	accountId = accountId & 0xffffffff
	index = accountId >> 27
	userId = accountId & 0x7ffffff
	return
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
