package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

var dlrs *redis.Client

func init() {
	options := redis.Options{
		Addr:     "192.168.131.125:6379",
		Password: "abc123",
	}
	dlrs = redis.NewClient(&options)
}

func main() {

	key := "dfasdfa"
	requestId := "dafsdfasdf"
	fmt.Println("lock --> ", GetDistributeLock(dlrs, key, requestId, 10*time.Second))
	defer ReleaseDistributeLock(dlrs, key, requestId)
	time.Sleep(3 * time.Second)
}

func GetDistributeLock(client *redis.Client, lockKey, requestId string, expire time.Duration) bool {
	return client.SetNX(lockKey, requestId, expire).Val()
}

func ReleaseDistributeLock(client *redis.Client, lockKey, requestId string) bool {
	script := "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end"
	val := client.Eval(script, []string{lockKey}, requestId).Val()
	fmt.Println("fdafsadfasdfa")
	return val.(int64) == 1
}
