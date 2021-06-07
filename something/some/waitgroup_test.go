package some

import (
	"fmt"
	"sync"
	"testing"
)

//waitGroup相关demo
func TestWaitGroup(t *testing.T) {

	list := make([]int, 0, 5)
	list[3] = 1
	fmt.Println(list)
	//sync.Mutex{}

	var wait sync.WaitGroup
	/*count := 10
	wait.Add(count)
	for i := 0; i < count; i++ {
		go func() {

		}()
	}*/
	fmt.Println("start wait")
	wait.Wait()
	fmt.Println("wait done")
	copyDemo(wait)
	var w = wait
	w.Wait()
	fmt.Println("owieur")



}

func copyDemo(w sync.WaitGroup) {
	w.Wait()
	fmt.Println("ldkaj")
}
