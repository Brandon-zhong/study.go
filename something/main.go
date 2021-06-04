package main

import (
	"fmt"
	"strconv"
)

func OrderBySign(input []int) []int {
	//
	// 请您补充完整这个函数，实现题目要求。
	//
	first, zeroIndex := 0, 0
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			zeroIndex = i
			continue
		}
		if input[i] > 0 {
			if i != first {
				input[first], input[i] = input[i], input[first]
			}
			first++
		}
	}
	if first < len(input) && zeroIndex != 0 {
		input[first], input[zeroIndex] = input[zeroIndex], input[first]
	}
	return input
}

// 请不要修改以下代码
func main() {
	input1 := []int{6, 4, -3, 5, -2, -1, 0, 1, -9}
	input2 := []int{1, 2, 3}
	input3 := []int{0, 0, 1, -1}
	input4 := []int{1}
	input5 := []int{-1, 0, 0, 1, -1}

	fmt.Print("input1 result: ")
	fmt.Println(OrderBySign(input1))
	fmt.Print("input2 result: ")
	fmt.Println(OrderBySign(input2))
	fmt.Print("input3 result: ")
	fmt.Println(OrderBySign(input3))
	fmt.Print("input4 result: ")
	fmt.Println(OrderBySign(input4))
	fmt.Print("input5 result: ")
	fmt.Println(OrderBySign(input5))
}
func demo(list []string) {
	list[0] = "hahah"
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
