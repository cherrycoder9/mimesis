package commands

import (
	"command/receiver"
	"fmt"
)

// 주문 취소를 캡슐화하는 명령 객체
type CancelCommand struct {
	Barista *receiver.Barista
	Menu    string
}

// 주문 취소 명령 실행
func (c *CancelCommand) Execute() {
	fmt.Println("[명령 실행] 주문을 취소하는 중입니다...")
	c.Barista.CancelCoffee(c.Menu)
}
