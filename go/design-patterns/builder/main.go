package main

import (
	"builder/coffee"
	"fmt"
)

// 빌더 패턴은 복잡한 객체를 단계별로 구성하는 방법
// 매개변수가 많거나 선택적 속성이 많은 객체를 만들때 효과적
func main() {
	fmt.Println("커피 주문 생성 시작")

	// CoffeeBuilder를 생성해 단계별로 커피 옵션 설정
	builder := coffee.NewCoffeeBuilder()

	// 빌더 체인을 이용해 주문 옵션 설정 후 최종 커피 주문 객체 생성
	myCoffee := builder.
		SetSize("대").
		SetCoffeeType("아메리카노").
		SetMilk("두유").
		AddShots(2).
		SetSyrup("카라멜").
		SetTemperature("따뜻한").
		Build()

	// 주문한 커피 정보 출력
	fmt.Println("[주문한 커피 정보]")
	fmt.Println(myCoffee)
}
