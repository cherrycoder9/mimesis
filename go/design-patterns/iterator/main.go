package main

import (
	"fmt"
	"iterator/menu"
)

// # 이터레이터 패턴 사용 이유
// 컬렉션 내부 구조를 노출하지 않고 요소들을 순차적으로 접근할 수 있음
// 서로 다른 자료구조를 동일한 방식으로 처리할 수 있게 해 코드의 재사용성 높임
// 내부 구현의 변경이 외부 코드에 영향을 주지 않도록 추상화 가능
func main() {
	fmt.Println("음료 메뉴")

	drinkMenu := menu.NewDrinkMenu()
	iterator := drinkMenu.CreateIterator()

	// 이터레이터를 통해 메뉴 항목을 순회하며 출력
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Printf("- %s (%.0f원): %s\n", item.Name, item.Price, item.Description)
	}
}
