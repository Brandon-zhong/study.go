package main

import "fmt"

func main() {

	var arr [3]int = [3]int{}
	for i := range arr {
		arr[i] = i + 5
	}
	fmt.Println(arr)
	var slice []int = arr[:]
	fmt.Println(slice)
	ints := appendInt(slice, 999)

	fmt.Println(ints)

	fmt.Println(cap(slice), " --- ", cap(ints))

}

func appendInt(x []int, y int) []int {

	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		//容量不够，创建新的数组
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
