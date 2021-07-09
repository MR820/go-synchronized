/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/7/9
 * Time 上午10:26
 */

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	var sm sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := atomic.AddInt64(&counter, 1)
			sm.Store(n, n*n)
		}()
	}
	wg.Wait()
	if v, ok := sm.Load(int64(1)); ok {
		fmt.Println(v)
	}

	sm.Range(func(k, v interface{}) bool {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
		return true
	})

}
