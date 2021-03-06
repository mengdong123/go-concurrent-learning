package main

import (
	"fmt"
	"sync"
)

// 在这个例子中，我们创建了 10 个 goroutine，同时不断地对一个变量（count）进行加1
// 操作，每个 goroutine 负责执行 10 万次的加 1 操作，我们期望的最后计数的结果是 10 *
// 100000 = 1000000 (一百万)。
// 每次运行，你都可能得到不同的结果，基本上不会得到理想中的一百万的结果。
// 使用go run -race mutex_01.go来查看data race的警告信息
func main() {
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Printf("count：%d\n", count)
}
