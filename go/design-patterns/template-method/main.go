package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"template-method/beverage"
)

// # 템플릿 메서드 패턴
// 작업 전체 흐름을 정의하면서 일부 단계를 하위 클래스나 구현체에 위임하는 방식
// 공통적인 프로세스를 재사용하면서 세부 구현만 다르게 해야 할때 효과적
// 코드 중복을 줄이고 일관성을 유지하지만 유연성이 부족해 확장이 어려울 수 있음
// 프로세스가 자주 바뀌거나 커스터마이징 범위가 넓은 경우 적합하지 않음
func main() {
	fmt.Println("어떤 음료를 주문하시겠습니까? (커피/차)")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "커피":
		coffee := &beverage.Coffee{}
		beverage.PrepareRecipe(coffee)
	case "차":
		tea := &beverage.Tea{}
		beverage.PrepareRecipe(tea)
	default:
		fmt.Println("잘못된 선택입니다. 프로그램을 종료합니다.")
	}
}
