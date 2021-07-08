/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/6/25
 * Time 下午3:37
 */

package main

import (
	"sync"
	"sync/atomic"
)

var counter int64

func incr() {
	for i := 0; i < 10; i++ {
		atomic.AddInt64(&counter, 1)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
	println(counter)
}
