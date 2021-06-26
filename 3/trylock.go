/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/6/25
 * Time 下午3:48
 */

package main

import "sync"

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:

	}
	return lockResult
}

func (l Lock) UnLock() {
	l.c <- struct{}{}
}

var counter int

func main() {
	var wg sync.WaitGroup
	var l = NewLock()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.UnLock()
		}()
	}
	wg.Wait()
	println(counter)
}
