package main

import (
	"adapter/adapter"
	"adapter/legacy"
	"fmt"
)

// 클라이언트가 기대하는 인터페이스와 기존 클래스의 인터페이스가 다를 때
// 두 인터페이스를 호환시키기 위해 사용됨

// Payment 인터페이스는 최신 시스템에서 사용되는 인터페이스
// 클라이언트는 이것으로 결제 기능 호출
type Payment interface {
	ProcessPayment(amount float64) string
}

func main() {
	var amount float64
	fmt.Print("결제 금액을 입력하세요: ")
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("입력 오류 발생:", err)
		return
	}

	// 레거시 결제 시스템은 별도 인터페이스를 가지고 있음
	// 최신 Payment 인터페이스에 맞추기 위해 어댑터를 사용
	// 어댑터는 레거시 시스템을 감싸서 최근 방식에서도 동작하게 함
	legacySystem :=
		legacy.LegacyPayment{} // 레거시 시스템 객체 생성
	paymentAdapter :=
		adapter.LegacyPaymentAdapter{Legacy: legacySystem} // 어댑터를 통한 래핑

	// 어댑터를 이용해 결제 처리
	result := paymentAdapter.ProcessPayment(amount)
	fmt.Println(result)
}
