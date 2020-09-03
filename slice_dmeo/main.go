package main

import "fmt"

func main() {

	arrayA := [2]int{100, 200}
	var arrayB [2]int

	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	testArray(arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	sli := arrayA[:]
	fmt.Printf("sli: %p, %p\n", sli, &sli)


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
