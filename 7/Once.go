/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/7/9
 * Time 上午10:15
 */

package main

import (
	"fmt"
	"sync"
)

var counter int

func incr() {
	for i := 0; i < 10; i++ {
		counter++
	}
}

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(incr)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
