package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"strconv"
)

var RS *redis.Client

func init() {
	options := redis.Options{
		Addr:     "192.168.131.125:6379",
		Password: "abc123",
	}
	RS = redis.NewClient(&options)
}

func main() {

	//RS.ZRangeByScore()
	//addData()
	key := "demo"
	lastCommentId := 17
	min := "(" + strconv.Itoa(lastCommentId)
	if lastCommentId == 0 {
		min = "+inf"
	}

	result, _ := RS.ZRevRangeByScore(key, &redis.ZRangeBy{Min: "-inf", Max: min, Offset: 0, Count: 4}).Result()
	for _, s := range result {
		fmt.Println(s)
	}

}

func addData() {
	RS.Del("demo")
	for i := 1; i < 25; i++ {
		RS.ZAdd("demo", &redis.Z{Member: i, Score: float64(i)})
	}
}
