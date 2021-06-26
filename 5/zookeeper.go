/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/6/26
 * Time 下午4:14
 */

package main

import (
	"sync"
	"time"

	"github.com/go-zookeeper/zk"
)

var counter int
var c *zk.Conn

func init() {
	conn, _, err := zk.Connect([]string{"node01", "node02", "node03", "node04"}, time.Second*10)
	if err != nil {
		panic(err)
	}
	c = conn
}

func incr() {

	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	err := l.Lock()
	if err != nil {
		panic(err)
	}
	counter++
	println("lock succ,do your business logic")
	l.Unlock()
	println("unlock succ,finish business logic")
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
