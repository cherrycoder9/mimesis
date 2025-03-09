package order

import "fmt"

// 상태가 수행해야 하는 동작 정의
type State interface {
	ProcessOrder(o *Order) // 주문 처리
	GetName() string       // 상태 이름 반환
}

// 주문이 접수된 상태
type ReceivedState struct{}

func (s *ReceivedState) ProcessOrder(o *Order) {
	fmt.Println("주문 접수 완료 -> 음료 준비 시작합니다.")
	o.SetState(&PreparingState{}) // 다음 상태로 변경
}

func (s *ReceivedState) GetName() string {
	return "주문 접수됨"
}

// 음료 준비중인 상태
type PreparingState struct{}

func (s *PreparingState) ProcessOrder(o *Order) {
	fmt.Println("음료 준비 완료 -> 픽업 준비합니다.")
	o.SetState(&ReadyState{}) // 다음 상태로 변경
}

func (s *PreparingState) GetName() string {
	return "음료 준비중"
}

// 픽업 가능한 상태
type ReadyState struct{}

func (s *ReadyState) ProcessOrder(o *Order) {
	fmt.Println("음료가 픽업 가능합니다. 감사합니다.")
	o.SetState(&CompletedState{}) // 다음 상태로 변경
}

func (s *ReadyState) GetName() string {
	return "픽업 준비 완료"
}

// 완료된 상태
type CompletedState struct{}

func (s *CompletedState) ProcessOrder(o *Order) {
	fmt.Println("주문이 이미 완료되었습니다. 추가 작업이 없습니다.")
}

func (s *CompletedState) GetName() string {
	return "완료됨"
}
