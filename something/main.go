package main

import (
	"fmt"
	"sync"
)

type A struct {
	id int
}

func main() {
	channel := make(chan int, 5)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for a := range channel {
			wg.Add(1)
			go func() {
				defer wg.Done()
				fmt.Println(a)
			}()
		}

	}()

	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)

	wg.Wait()
}
