package main

import (
	"fmt"
	"sync"
	"time"
)

func loadInteger(number int, wgLoadInteger *sync.WaitGroup) {
	defer wgLoadInteger.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("정수: %d\n", number+i)
	}
}

func loadFloat(number float64, wgLoadFloat *sync.WaitGroup) {
	defer wgLoadFloat.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("실수: %.1f\n", number+float64(i))
	}

}

func main() {
	var wgLoadInteger sync.WaitGroup
	var wgLoadFloat sync.WaitGroup
	wgLoadInteger.Add(3)
	wgLoadFloat.Add(3)

	fmt.Println("integer 불러오기 준비중")
	go loadInteger(1, &wgLoadInteger)
	go loadInteger(11, &wgLoadInteger)
	go loadInteger(21, &wgLoadInteger)
	wgLoadInteger.Wait()
	fmt.Println("integer 불러오기 완료")

	fmt.Println("float 불러오기 준비중")
	go loadFloat(1.0, &wgLoadFloat)
	go loadFloat(11.0, &wgLoadFloat)
	go loadFloat(21.0, &wgLoadFloat)
	wgLoadFloat.Wait()
	fmt.Println("float 불러오기 완료")
}
