package legacy

import "fmt"

// LegacyPayment는 기존 시스템에서 사용되는 결제 구조체
// 이 시스템은 결제 금액을 정수로 처리함
type LegacyPayment struct{}

// MakePayment 메서드는 정수형 금액을 받아 결제 수행
// 이 함수는 레거시 시스템의 인터페이스로 새로운 시스템과 사용방법이 다름
func (lp LegacyPayment) MakePayment(money int) string {
	return fmt.Sprintf("레거시 시스템: %d원 결제 완료", money)
}
