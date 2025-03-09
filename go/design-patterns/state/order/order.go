package order

import "fmt"

// 현재 상태를 관리
type Order struct {
	state State // 현재 주문의 상태
}

// 새 주문 객체 생성시 초기 상태는 ReceivedState로 설정
func NewOrder() *Order {
	return &Order{state: &ReceivedState{}}
}

// 상태 변경 메서드
func (o *Order) SetState(s State) {
	o.state = s
}

// 현재 상태를 진행시키는 메서드
func (o *Order) Proceed() {
	o.state.ProcessOrder(o)
}

// 현재 상태 이름을 반환하는 메서드
func (o *Order) CurrentState() {
	fmt.Printf("현재 주문 상태: [%s]\n", o.state.GetName())
}
