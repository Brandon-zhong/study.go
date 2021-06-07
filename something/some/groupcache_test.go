package some

import (
	"fmt"
	"github.com/golang/groupcache/singleflight"
	"sync"
	"testing"
	"time"
)

//减少并发时，多个协程重复调用造成的资源开销
func TestSingleFlight(t *testing.T) {

	wg := singleflight.Group{}
	var w sync.WaitGroup
	var w2 sync.WaitGroup
	w.Add(1)
	w2.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			w2.Done()
			w.Wait()
			result, _ := wg.Do("kdlajflkdsj", func() (interface{}, error) {
				fmt.Println("itpequrowiej===========")
				return "hahah", nil
			})
			fmt.Println(result)
		}()
	}
	w2.Wait()
	w.Done()
	time.Sleep(5 * time.Second)
}
