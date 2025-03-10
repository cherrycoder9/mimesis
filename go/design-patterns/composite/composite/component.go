package composite

import "fmt"

// 단일 메뉴 항목(Leaf)과 복합 메뉴 항목(Composite) 모두가
// 구현해야 하는 메서드를 정의하는 인터페이스
// 클라이언트는 메뉴의 구성 요소를 동일한 방식으로 다룰 수 있음
type MenuComponent interface {
	// 메뉴 항목의 정보를 들여쓰기를 적용해 출력하는 메서드
	Print(indent string)
}

// Leaf는 자식 메뉴를 가질 필요가 없으므로 Add 메서드는 구현하지 않으며
// Composite는 자식 메뉴를 관리하기 위한 Add 메서드를 별도로 제공함

func dummy() {
	fmt.Println("MenuComponent 인터페이스 예제")
}
