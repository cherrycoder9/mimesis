package main

import (
	"decorator/beverage"
	"decorator/decorator"
	"fmt"
)

// 데코레이터 패턴은 객체에 동적으로 새로운 책임이나 기능을 추가할 수 있게 해주는 패턴
// 상속 대신 조합을 활용함, 기본 객체 코드를 변경하지 않고 기능을 확장할때 사용
// Go에서는 인터페이스를 구현한 구조체를 감싸는 방식으로 적용할수 있음
func main() {
	// 기본 음료(에스프레스) 객체 생성
	var myBeverage beverage.Beverage = beverage.NewEspresso()
	fmt.Println("주문한 기본 음료:", myBeverage.GetDescription())
	fmt.Printf("기본 음료 가격: %.2f원\n", myBeverage.Cost())

	// 우유 데코레이터를 적용해 음료에 우유를 추가
	myBeverage = decorator.NewMilkDecorator(myBeverage)
	fmt.Println("우유 추가 후 음료:", myBeverage.GetDescription())
	fmt.Printf("우유 추가 후 가격: %.2f원\n", myBeverage.Cost())

	// 설탕 데코레이터를 적용해 음료에 설탕 추가
	myBeverage = decorator.NewSugarDecorator(myBeverage)
	fmt.Println("설탕 추가 후 음료:", myBeverage.GetDescription())
	fmt.Printf("최종 음료 가격: %.2f원\n", myBeverage.Cost())
}
