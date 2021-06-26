/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/6/25
 * Time 下午4:00
 */

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	c := redis.NewClient(&redis.Options{
		Addr:     "node01:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	client = c
}

func incr() {

	var lockKey = "counter_lock"
	var counterKey = "counter"

	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		//incr()
		return
	}
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			panic("set value error!")
		}
	}
	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
	} else {
		panic("unlock failed")

	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
	getResp := client.Get("counter")
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		println(cntValue)
	}
}
