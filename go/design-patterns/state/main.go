package main

import (
	"state/order"
)

// # 왜 상태 패턴을 쓰는가?
// 상태 패턴은 객체가 상태에 따라 행동을 변경할 수 있도록 하는 패턴
// 객체 상태 변화에 따라 특정 행위가 바뀌어야 할때, 상태 패턴을 통해 상태마다의 동작을 독립된 클래스로 캡슐화해 관리하면, 조건문을 많이 쓰는 복잡한 코드를 방지하고 유연한 코드 유지보수 가능
func main() {
	// 새 주문을 생성하고 초기 상태 출력
	order := order.NewOrder()
	order.CurrentState()

	// 주문 상태를 진행시키면서 상태 변화 출력
	order.Proceed()
	order.CurrentState()

	order.Proceed()
	order.CurrentState()

	order.Proceed()
	order.CurrentState()

	// 완료된 주문에서 추가 Proceed 호출시의 반응
	order.Proceed()
}
