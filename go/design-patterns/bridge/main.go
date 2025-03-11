package main

import (
	"bridge/abstraction"
	"bridge/implementation"
	"fmt"
)

// # 브리지 패턴
// 추상화와 구현을 분리해 독립적으로 확장할 수 있게 함
// 구현 세부 사항이 자주 바뀌거나 플랫폼별로 달라질때 유용함
// ex: 그래픽 라이브러리에서 렌더링 방식을 분리할 수 있음
// 추상화는 클라이언트와 연결되고 구현을 별도로 관리됨
func main() {
	// 에스프레소 머신을 사용해 커피 추출 방식 구현
	espressoMachine := implementation.NewEspressoMachine()
	// 아메리카노 커피는 에스프레소 머신을 이용해 추출됨
	americano := abstraction.NewCoffee("아메리카노", espressoMachine)
	fmt.Println(americano.Serve())

	// 핸드드립 방식을 사용해 커피 추출 방식 구현
	handBrew := implementation.NewHandBrew()
	// "라떼" 커피는 핸드드립 방식을 이용해 추출됨
	latte := abstraction.NewCoffee("라떼", handBrew)
	fmt.Println(latte.Serve())
}
