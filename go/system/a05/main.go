package main

import (
	"fmt"
	"sync"
)

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
