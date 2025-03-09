package mediator

import "fmt"

// 바리스타 인터페이스
type Barista interface {
	MakeDrink(drink string, customer string)
	GetName() string
}

// 구체적인 바리스타 구현
type CafeBeneBarista struct {
	Name string
}

// 음료를 제조하는 메서드 (실제 바리스타의 행위)
func (b *CafeBeneBarista) MakeDrink(drink string, customer string) {
	fmt.Printf("[바리스타 %s] %s님을 위한 %s 준비 완료! ☕️\n", b.Name, customer, drink)
}

// 바리스타의 이름을 반환하는 메서드
func (b *CafeBeneBarista) GetName() string {
	return b.Name
}
