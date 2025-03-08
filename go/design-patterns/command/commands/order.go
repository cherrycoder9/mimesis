package commands

import (
	"command/receiver"
	"fmt"
)

// 커피 주문을 캡슐화하는 명령 객체
type OrderCommand struct {
	Barista *receiver.Barista
	Menu    string
}

// 주문 명령을 실행
func (o *OrderCommand) Execute() {
	fmt.Println("[명령 실행] 주문을 처리 중입니다...")
	o.Barista.MakeCoffee(o.Menu)
}
