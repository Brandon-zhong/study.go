package main

import (
	"fmt"
	"strconv"
	"time"
)

//注意，select只会执行一次，不会反复执行
func main() {
	demo1()
}

func demo() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
		time.Sleep(5 * time.Second)
	}()

	go func() {
		chan2 <- 2
		time.Sleep(5 * time.Second)
	}()

	for {
		select {
		case <-chan1:
			fmt.Println("chan1 ready.")
			goto p
		case <-chan2:
			fmt.Println("chan2 ready.")
			goto p
		default:
			fmt.Println("default")
		}
	}
p:
	fmt.Println("main.exit")

}

//select随机检测各case语句中的channel是否ready,如果某个case中的channel已经ready，则执行相应的case然后退出select流程
//如果所有的channel都不为ready且没有default的话，则阻塞等待各channel
func demo1() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	writeFlag := true
	go func() {
		for {
			if writeFlag {
				chan1 <- 1
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			if writeFlag {
				chan2 <- 2
			}
			time.Sleep(time.Second)
		}
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}
	fmt.Println("main.exit")
}

//测试select的随机检测各case, 准备两个有元素的管道，然后select这两个管道并打印，查看打印的顺序
//从打印结果可知，select不是按照case的顺序来检测channel是否ready，而是随机检测channel是否ready
func demo3() {

	chan1 := make(chan string, 10)
	chan2 := make(chan string, 10)
	for i := 0; i < 10; i++ {
		chan1 <- "chan1 " + strconv.Itoa(i)
	}
	for i := 0; i < 10; i++ {
		chan2 <- "chan2 " + strconv.Itoa(i)
	}

	for i := 0; i < 20; i++ {
		select {
		case s := <-chan1:
			fmt.Println(s)
		case s := <-chan2:
			fmt.Println(s)
		}
		fmt.Println("----")
		time.Sleep(time.Second)
	}
}

//已关闭的channel是可以读的，此方法会随机执行两个case的打印代码
func demo4() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		close(chan1)
	}()
	go func() {
		close(chan2)
	}()

	for {
		select {
		case <-chan1:
			fmt.Println("chan1 ready.")
		case <-chan2:
			fmt.Println("chan2 ready.")
		}
		time.Sleep(time.Second)
	}
	//fmt.Println("main.exit")
}

//空的select没有case可以检测，所以当前协程会阻塞
func demo5() {
	select {}
}
