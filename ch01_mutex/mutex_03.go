package main

import (
	"fmt"
	"sync"
)

// 采用嵌入字段的方式直接调用mutex
type Counter struct {
	sync.Mutex
	Count uint64
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Print(counter)
}
