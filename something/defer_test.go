package main

import (
	"fmt"
	"testing"
)
// 输出 1 2 即使是panic，对defer的触发也是遵循先进后出的原则
func TestDefer(t *testing.T){
	defer func() {
		if err := recover(); err !=nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		fmt.Println(1)
	}()
	panic(2)
	panic(3)
}
