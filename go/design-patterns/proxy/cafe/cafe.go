package cafe

import "fmt"

// 음료 주문을 처리하는 인터페이스
type Menu interface {
	Order(drink string)
}

// 실제 매장(리소스가 많은 원본 객체)에 대응하는 구조체
type CafeBene struct{}

// Order는 실제 주문을 처리하는 메서드
func (s *CafeBene) Order(drink string) {
	fmt.Printf("✅ [%s] 주문이 정상적으로 접수되었습니다.\n", drink)
}
