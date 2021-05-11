package main

import (
	"time"
)

func main() {

}

//超时控制，某些操作在指定的时间内完成，超时退出不再等待
func timeoutControlDemo() {
	select {
	case <-time.After(time.Second * 3):
	//case  _, _ := http.Get("www.google.com"):

	}
}

func cronDemo() {

}


