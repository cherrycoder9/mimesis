package main

import "fmt"

// Flyweight 인터페이스
type Beverage interface {
	Serve(size string)
}

// Flyweight 객체로 공유되는 실제 음료
type ConcreteBeverage struct {
	name        string // 음료 이름 (공유 데이터)
	description string // 음료 설명 (공유 데이터)
}

// 음료를 손님에게 제공하는 함수
func (b *ConcreteBeverage) Serve(size string) {
	fmt.Printf("[서빙] %s(%s) 음료 나왔습니다! - 설명: %s\n", b.name, size, b.description)
}
