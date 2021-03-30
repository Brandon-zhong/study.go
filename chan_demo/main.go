package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	num := 10
	var ch = make(chan int)
	/*for i := 0; i < num; i++ {
		ch <- i
	}*/
	close(ch)
	for i := 0; i < num+5; i++ {
		fmt.Println(<-ch)
	}

}

func demo1() {
	channel := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for a := range channel {
					fmt.Println(a)
					time.Sleep(100 * time.Millisecond)
				}
			}()
		}
	}()

	for i := 0; i < 11; i++ {
		channel <- i
	}
	close(channel)

	wg.Wait()
	time.Sleep(10 * time.Second)
	fmt.Println("all done.")
}
