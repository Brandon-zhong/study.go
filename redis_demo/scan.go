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

// func main() {

// 	scan()

// }

func scan() {
	fmt.Println("start match --> ", match)
	cursor := uint64(0)
again:
	result, cursor, err := rs.Scan(cursor, match, 1000).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(len(result),"--", cursor)
	for _, str := range result {
		fmt.Println(str)
	}
	if cursor != 0 {
		time.Sleep(100 * time.Millisecond)
		goto again
	}
}

func insertDemoData() {
	for i := 0; i < 100000; i++ {
		t := strconv.Itoa(i)
		rs.Set("demo_key_"+t, t, 0)
	}
}
