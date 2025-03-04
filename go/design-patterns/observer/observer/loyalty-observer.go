package observer

import "fmt"

// 포인트 적립 시스템을 모방한 옵저버
type LoyaltyObserver struct{}

// LoyaltyObserver 객체를 생성하는 생성자
func NewLoyaltyObserver() *LoyaltyObserver {
	return &LoyaltyObserver{}
}

// 포인트 적립 시스템이 주문 알림을 받았을 때 수행할 작업 구현
func (lo *LoyaltyObserver) Update(message string) {
	fmt.Println("[포인트 시스템] 주문 알림 받음:", message)
	fmt.Println("[포인트 시스템] 고객 포인트 적립 진행 중...")
}
