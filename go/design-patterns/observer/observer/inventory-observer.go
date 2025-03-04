package observer

import "fmt"

// 재고 관리 시스템을 모방한 옵저버
type InventoryObserver struct{}

// InventoryObserver 객체를 생성하는 생성자
func NewInventoryObserver() *InventoryObserver {
	return &InventoryObserver{}
}

// 재고 관리 시스템이 주문 알림을 받았을때 수행할 작업을 구현
func (io *InventoryObserver) Update(message string) {
	fmt.Println("[재고 시스템] 주문 알림 받음:", message)
	fmt.Println("[재고 시스템] 재고 업데이트 진행 중...")
}
