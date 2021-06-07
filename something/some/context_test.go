package some

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//context 包相关函数测试
func TestChancelCtx(t *testing.T) {
	//context可以解决退出所有派生协程的操作，但是普通的chan也可以做到，不是必须的
	ctx, cancel := context.WithCancel(context.Background())
	go HandelRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines!")
	cancel()
	//Just for test whether sub goroutines exit or not
	time.Sleep(5 * time.Second)

	var fun = func(ch <-chan interface{}, name string) {
		for {
			select {
			case <-ch:
				fmt.Println(name, "is done!")
				return
			default:
				fmt.Println(name, "is running")
				time.Sleep(time.Second)
			}
		}
	}
	fmt.Println("==============================")
	//下面是用普通的chan来实现退出所有派生协程的操作
	var done = make(chan interface{})
	go func(ch <-chan interface{}, name string) {
		go fun(done, "redis")
		go fun(done, "mysql")
		for {
			select {
			case <-ch:
				fmt.Println(name, "is done!")
				return
			default:
				fmt.Println(name, "is running")
				time.Sleep(time.Second)
			}
		}
	}(done, "handler")
	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all goroutines")
	close(done)
	time.Sleep(5 * time.Second)

}

func HandelRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)  }
	}
}

func WriteRedis(ctx context.Context) {
	for {
		select {  case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(2 * time.Second)  }
	}
}
func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(2 * time.Second)  }
	}
}
