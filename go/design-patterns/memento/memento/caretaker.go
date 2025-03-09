package memento

// 메멘토들을 관리(보관)하는 역할을 함
type Caretaker struct {
	mementos []*Memento
}

// 메멘토 추가
func (c *Caretaker) Add(m *Memento) {
	c.mementos = append(c.mementos, m)
}

// 특정 메멘토 가져오기
func (c *Caretaker) Get(index int) *Memento {
	return c.mementos[index]
}
