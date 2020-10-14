package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	/*if from, err := readFrom(os.Stdin, 10); err == nil || err == io.EOF {
		fmt.Println(string(from), "---")
	}
	println("fjldkasjfl")*/

	file, err := os.Open("C:\\Users\\iyingdi\\Downloads\\reader.go")
	if err != nil {
		fmt.Println("open file fail")
		return
	}

	bytes := make([]byte, 2048)
	from, err := readFrom(file, 1024)
	if err == nil {

	} else if err == io.EOF {

	}

}

func readFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	fmt.Println("n --> ", n)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}
