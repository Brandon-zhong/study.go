package main

import "fmt"

func GetCount(str string) (count int) {
	var vowelsMap = map[uint8]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	for i := 0; i < len(str); i++ {
		if _, ok := vowelsMap[str[i]]; ok {
			count++
		}
	}
	// Enter solution here
	return count
}

func main() {
	fmt.Println(GetCount("abracadabra"))
}
