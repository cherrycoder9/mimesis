package main

import (
	"fmt"
	"observer/observer"
	"observer/subject"
)

// 객체의 상태 변화가 있을때 이에 의존하는 다른 객체들에게 자동으로 알리고 업데이트하는 방식
// 주로 1대다 관계를 관리함, 이벤트 시스템이나 GUI에서 상태 변화를 실시간으로 반영할때 유용
// 느슨한 결합을 유지해 확장성을 높이지만 관찰자가 많아지면 성능 저하나 메모리 누수 발생 가능성
// 단순한 상태 공유나 빈번한 업데이트가 필요 없는 경우엔 사용하지 않는게 나음
func main() {
	// 주문을 관리하는 주체(Subject)를 생성
	orderSubject := subject.NewOrderSubject()

	// 옵저버 객체 생성 (재고 관리 시스템과 포인트 적립 시스템)
	inventoryObserver := observer.NewInventoryObserver()
	loyaltyObserver := observer.NewLoyaltyObserver()

	// 생성된 옵저버들을 Subject에 등록함
	orderSubject.Registers(inventoryObserver)
	orderSubject.Registers(loyaltyObserver)

	// 새로운 커피 주문 발생
	fmt.Println("새로운 커피 주문: 카푸치노")
	orderSubject.NewOrder("카푸치노")

}
