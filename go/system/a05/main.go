package main

import (
	"fmt"
	"sync"
)

// 아래 명령을 Go 런타임에 레이스 감지기를 사입해
// 공유 메모리에서의 비정상적인 접근을 실시간으로 추적할 수 있음
// go run -race main.go
// go test -race

func main() {
	fmt.Println("Total Items Packed:", PackItems(0))
}

func PackItems(totalItems int) int {
	const workers = 2
	const itemsPerWorker = 1000
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < itemsPerWorker; j++ {
				mu.Lock()
				totalItems++
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	return totalItems
}
