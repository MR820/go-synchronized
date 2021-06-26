/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/6/25
 * Time 下午3:37
 */

package main

import "sync"

var counter int

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	println(counter)
}
