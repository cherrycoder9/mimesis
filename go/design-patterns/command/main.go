package main

import (
	"command/commands"
	"command/invoker"
	"command/receiver"
	"fmt"
)

// # 커맨드 패턴 사용이유
// 실행할 작업을 객체로 캡슐화해 명령을 내리는(invoker) 객체와 명령을 실제 수행하는(receiver) 객체를 분리할 수 있음
// 요청을 큐로 관리하거나, 로그 기록 및 취소 기능을 쉽게 추가할 수 있음
// 확장성을 높이고 유연한 시스템 설계가 가능함
func main() {
	// 바리스타 생성 (Receiver)
	barista := &receiver.Barista{}

	// 캐셔 생성 (Invoker)
	cashier := &invoker.Cashier{}

	fmt.Println("주문 시스템 시작")

	// 명령 객체 생성
	orderAmericano := &commands.OrderCommand{
		Barista: barista,
		Menu:    "아메리카노",
	}

	orderLatte := &commands.OrderCommand{
		Barista: barista,
		Menu:    "카페라떼",
	}

	cancelLatte := &commands.CancelCommand{
		Barista: barista,
		Menu:    "카페라떼",
	}

	// 주문받기
	cashier.TakeOrder(orderAmericano)
	cashier.TakeOrder(orderLatte)
	cashier.TakeOrder(cancelLatte)

	// 주문 실행
	cashier.ExecuteOrders()
}
