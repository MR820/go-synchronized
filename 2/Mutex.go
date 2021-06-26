/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/6/25
 * Time 下午3:42
 */

package main

import "sync"

var counter int
var l sync.Mutex

func incr() {
	l.Lock()
	counter++
	l.Unlock()
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
	println(counter)
}
