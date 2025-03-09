package mediator

import "fmt"

// 중재자 인터페이스, 객체 간 통신을 중앙에서 제어
type Mediator interface {
	RegisterBarista(Barista)
	SendOrder(drink string, customer string)
}

// 구체적인 중재자 구조체
type CoffeeMediator struct {
	baristas []Barista
}

// 바리스타 등록 메서드
func (m *CoffeeMediator) RegisterBarista(b Barista) {
	m.baristas = append(m.baristas, b)
}

// 고객의 주문을 받아 바리스타에게 전달 (실제로는 주문을 분산처리하거나 적절히 선택 가능)
func (m *CoffeeMediator) SendOrder(drink string, customer string) {
	if len(m.baristas) == 0 {
		fmt.Println("바리스타가 없습니다.")
		return
	}

	// 단순로직, 첫번째 바리스타가 주문 처리
	m.baristas[0].MakeDrink(drink, customer)
}
