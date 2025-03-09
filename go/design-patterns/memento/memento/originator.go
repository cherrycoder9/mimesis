package memento

import "fmt"

// 현재의 주문 상태를 관리
type Originator struct {
	state string
}

// 현재 주문 상태를 설정
func (o *Originator) SetState(state string) {
	fmt.Println("주문 상태 변경:", state)
	o.state = state
}

// 현재 상태를 가져옴
func (o *Originator) GetState() string {
	return o.state
}

// 현재 상태로 메멘토 객체를 생성해 반환
func (o *Originator) SaveToMemento() *Memento {
	fmt.Println("상태 저장 중...")
	return NewMemento(o.state)
}

// 메멘토 객체에서 저장된 상태를 가져와서 복원
func (o *Originator) RestoreFromMemento(m *Memento) {
	o.state = m.GetState()
	fmt.Println("상태 복원됨:", o.state)
}
