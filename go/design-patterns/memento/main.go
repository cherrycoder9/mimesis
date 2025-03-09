package main

import (
	"fmt"
	"memento/memento"
)

// 메멘토 패턴은 객체 상태를 캡슐화하여 저장하고 필요할때 복구할 수 있게 해주는 디자인 패턴
// 객체의 상태를 직접 노출하지 않고도 복원할 수 있으며 되돌리기 기능 등을 구현할때 자주 사용
// 고객이 주문한 음료의 상태를 저장하고, 필요시 주문 상태를 이전 단계로 되돌리는것 구현
func main() {
	fmt.Println("스타벅스 주문 관리 시스템")

	// 객체 초기화
	caretaker := &memento.Caretaker{}
	order := &memento.Originator{}

	// 상태 변경 및 저장
	order.SetState("에스프레소 줌누")
	caretaker.Add(order.SaveToMemento())

	order.SetState("우유 추가")
	caretaker.Add(order.SaveToMemento())

	order.SetState("카라멜 시럽 추가")
	caretaker.Add(order.SaveToMemento())

	order.SetState("휘핑크림 추가")
	// 휘핑크림 추가후 상태 저장하지 않음
	fmt.Println("\n-- 현재 주문 상태 --")
	fmt.Println("✅", order.GetState())

	// 이전 상태로 되돌리기 (Undo 기능)
	fmt.Println("\n이전 상태로 되돌리는 중...")
	order.RestoreFromMemento(caretaker.Get(2)) // 카라멜 시럽 추가 상태
	fmt.Println("✅ 현재 상태: ", order.GetState())

	fmt.Println("\n한단계 더 되돌리는 중...")
	order.RestoreFromMemento(caretaker.Get(1)) // 우유 추가 상태
	fmt.Println("✅ 현재 상태:", order.GetState())
}
