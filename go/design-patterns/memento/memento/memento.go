package memento

// 이 구조체는 Originator의 상태를 저장하는 역할
type Memento struct {
	state string
}

// 새 메멘토 객체 생성
func NewMemento(state string) *Memento {
	return &Memento{
		state: state,
	}
}

// 저장된 상태를 반환
func (m *Memento) GetState() string {
	return m.state
}
