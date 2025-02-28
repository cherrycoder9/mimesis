package cafe

import "fmt"

// OrderDrink 함수는 팩토리 메서드를 이용해 음료 주문을 처리함
// 클라이언트(메인 함수)는 어떤 팩토리를 사용할지 결정하고
// 이 함수는 해당 팩토리를 통해 음료 객체를 생성한후 준비 과정 실행
func OrderDrink(factory DrinkFactory) {
	// 팩토리 메서드를 통해 원하는 음료 객체 생성
	drink := factory.CreateDrink()

	// 음료 준비 과정 실행
	fmt.Println("음료를 준비중입니다...")
	drink.Prepare()
}
