package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/duration"
	"time"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan int) {
	fmt.Printf("generate --> %p\n", ch)
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	fmt.Printf("filter in --> %p, out --> %p\n", in, out)
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func main() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for i := 0; i < 10; i++ {
		fmt.Printf("asdf --> %p\n", ch)
		prime := <-ch
		//fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}


func timeTick() {
	rate_per_sec := 10
	var dur time.Duration = 1e9
	tick := time.Tick(dur)


}
