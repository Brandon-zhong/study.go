package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	s := make([]int, 0)
	oldCap := cap(s)
	for i := 0; i < 2048 * 2 * 2 * 2 * 2 * 2; i++ {
		s = append(s, i)
		newCap := cap(s)
		if newCap != oldCap {
			fmt.Printf("[%d -> %d] cap = %d | after append %d cap = %d  difference = %d  rate=%-2.4f\n", 0, i-1, oldCap, i, newCap, newCap-oldCap, float64(newCap)/float64(oldCap))
			oldCap = newCap
		}

	}
}

func testArray(x [2]int) {
	x[0] = 666
	fmt.Printf("func Array : %p , %v\n", &x, x)
}

func testArray1() {
	var c int = 3
	var array customArray = make(customArray, c)
	for i := 0; i < c; i++ {
		array[i] = 3 + i
	}
	fmt.Printf("addr --> %p\n", array)
	fmt.Printf("addr --> %p\n", &array)
	fmt.Println(array)
	array.update(2, 10)
	array.update(0, 20)
	fmt.Println(array)

	update(array, 2, 666)
	fmt.Println(array)
}

type customArray []int

func (c customArray) update(index, value int) {
	c[index] = value
	fmt.Printf("addr --> %p\n", c)
	fmt.Printf("addr --> %p\n", &c)
}

func update(array customArray, index, value int) {
	array[index] = value
	fmt.Printf("addr --> %p\n", array)
	fmt.Printf("addr --> %p\n", &array)
}


func testSeeker() {
	reader := strings.NewReader("this is first go application")
	reader.Seek(-6, io.SeekEnd)
	readRune, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", readRune)
}
