package receiver

import "fmt"

// 커피를 만드는 직원 역할 (Receiver)
type Barista struct {
}

// 커피를 만들어주는 기능 수행
func (b *Barista) MakeCoffee(menu string) {
	fmt.Printf("바리스타: [%s] 커피를 만들었습니다.\n", menu)
}

// 커피 주문을 취소하는 기능 수행
func (b *Barista) CancelCoffee(menu string) {
	fmt.Printf("바리스타: [%s] 주문을 취소했습니다.\n", menu)
}
