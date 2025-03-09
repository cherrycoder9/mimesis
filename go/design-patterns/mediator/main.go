package main

import (
	"fmt"
	"mediator/mediator"
)

// # 중재자 패턴
// 여러 객체 간의 복잡한 의존성을 최소화하고, 객체들 간의 상호작용을 중앙에서 통제할 수 있도록 함. 특히 객체 수가 많아지거나 복잡한 비즈니스 로직이 얽혀 있을때, 서로 직접 통신하는 대신 중재자를 통해 관리하면 유지보수가 수월해짐
func main() {
	// 중재자 생성
	coffeeMediator := &mediator.CoffeeMediator{}

	// 바리스타 생성 및 중재자에 등록
	barista1 := &mediator.CafeBeneBarista{Name: "민철"}
	barista2 := &mediator.CafeBeneBarista{Name: "진호"}

	coffeeMediator.RegisterBarista(barista1)
	coffeeMediator.RegisterBarista(barista2)

	fmt.Println("☕️ 카페베네 중재자 시스템 시작 ☕️")

	// 고객 주문 처리
	coffeeMediator.SendOrder("카라멜 마키아또", "민수")
	coffeeMediator.SendOrder("아이스 아메리카노", "지영")
}
