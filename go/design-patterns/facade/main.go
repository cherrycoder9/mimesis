package main

import (
	"facade/facade"
	"fmt"
)

// 복잡한 서브시스템을 단순화된 인터페이스로 감싸서 클라이언트가 쉽게 사용할 수 있게 하는 패턴
// 여러 모듈이나 라이브러리를 통합해 일관된 접근을 제공할때 유용함
// 복잡한 API 호출 과정을 단일 함수로 줄일 수 있음
func main() {
	fmt.Println("카페베네 커피 주문 시스템에 오신 것을 환영합니다!")

	// CoffeeFacade 객체 생성 - 서브시스템들의 복잡한 인터페이스를 감춤
	coffeeFacade := facade.NewCoffeeFacade()

	// 아메리카노 주문 진행 (단일 메서드 호출로 모든 서브시스템이 작동)
	fmt.Println("아메리카노 주문 중...")
	coffeeFacade.OrderAmericano()

	fmt.Println("커피가 준비되었습니다. 즐거운 시간 보내세요!")
}
