package main

import (
	"bufio"
	"factory-method/cafe"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 팩토리 메서드 패턴을 활용해 카페에서 음료를 주문하는 예제
	// 이 패턴은 객체 생성 로직을 클라이언트 코드에서 분리해
	// 어떤 객체를 생성할지 결정하는 책임을 서브팩토리에게 위임함
	// 새로운 음료가 추가되더라도 클라이언트 코드를 수정할 필요가 없어짐

	// 사용자에게 주문할 음료를 입력받음
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("주문할 음료를 입력하세요 (커피, 차): ")
	orderInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("입력 오류:", err)
		return
	}
	orderInput = strings.TrimSpace(orderInput)

	// 사용자가 입력한 음료에 맞는 팩토리 객체 선택
	var factory cafe.DrinkFactory
	switch orderInput {
	case "커피":
		factory = &cafe.CoffeeFactory{}
	case "차":
		factory = &cafe.TeaFactory{}
	default:
		fmt.Println("알 수 없는 주문입니다.")
		return
	}

	fmt.Printf("%s 주문을 접수했습니다.\n", orderInput)
	// 선택된 팩토리를 사용해 음료 주문 처리
	cafe.OrderDrink(factory)
}
