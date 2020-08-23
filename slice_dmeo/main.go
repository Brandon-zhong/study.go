package main

import "fmt"

func main() {
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

}

type customArray []int

func (c customArray) update(index, value int) {
	c[index] = value
	fmt.Printf("addr --> %p\n", c)
	fmt.Printf("addr --> %p\n", &c)
}
