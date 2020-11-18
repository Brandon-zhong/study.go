package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis/v7"
	"os"
	"strconv"
	"time"
)

var rs *redis.Client
var params = &redis.Options{}
var match string

func init() {

	flag.StringVar(&params.Addr, "h", "", "remote redis host")
	flag.StringVar(&params.Password, "p", "", "redis password")
	flag.StringVar(&match, "match", "", "match keys")
	flag.Parse()
	checkParams()

	rs = redis.NewClient(params)
}

func checkParams() {
	if params.Addr == "" || params.Password == "" || match == "" {
		fmt.Println("params illegal, please check!")
		os.Exit(0)
	}
}

func main() {

	scan()

}

func scan() {
	cursor := uint64(0)
	match := "demo_key*"
	count := 0
	scanNum := 0

again:
	scanNum += 1
	result, cursor, err := rs.Scan(cursor, match, 1000).Result()
	if err != nil {
		return
	}
	count += len(result)
	if cursor != 0 {
		time.Sleep(100 * time.Millisecond)
		if scanNum%50 == 0 {
			fmt.Println(fmt.Sprintf("match --> %s , count --> %d", match, count))
		}
		goto again
	}
	fmt.Println(fmt.Sprintf("match --> %s , count --> %d , scanNum --> %d", match, count, scanNum))
}

func insertDemoData() {
	for i := 0; i < 100000; i++ {
		t := strconv.Itoa(i)
		rs.Set("demo_key_"+t, t, 0)
	}
}
