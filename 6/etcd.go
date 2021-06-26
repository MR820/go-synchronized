/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/6/26
 * Time 下午4:48
 */

package main

import (
	"log"
	"sync"

	"github.com/zieckey/etcdsync"
)

var counter int
var m *etcdsync.Mutex

func init() {
	mutex, err := etcdsync.New("/lock", 10, []string{"http://node01:2379"})
	if m == nil || err != nil {
		log.Printf("etcdsync.New failed")
		return
	}
	m = mutex
}

func incr() {
	err := m.Lock()
	if err != nil {
		log.Printf("etcdsync.Lock failed")
		return
	}

	counter++

	err = m.Unlock()
	if err != nil {
		log.Println("etcdsync.Unlock failed")
	} else {
		log.Printf("etcdsync.Unlock OK")
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
	println(counter)
}
