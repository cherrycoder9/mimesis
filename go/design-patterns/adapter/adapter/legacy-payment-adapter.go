package adapter

import (
	"adapter/legacy"
	"fmt"
)

// Payment 인터페이스와 호환되도록 어댑터 구현
type LegacyPaymentAdapter struct {
	Legacy legacy.LegacyPayment
}

// ProcessPayment 메서드는 float64 타입의 금액을 받아 int 타입으로 변환한 후 호출
func (adapter LegacyPaymentAdapter) ProcessPayment(amount float64) string {
	// float64 타입을 int로 변환 (소수점 이하를 버림)
	intAmount := int(amount)
	// 레거시 시스템의 결제 메서드 호출
	result := adapter.Legacy.MakePayment(intAmount)
	return fmt.Sprintf("어댑터를 통한 결제 결과: %s", result)
}
