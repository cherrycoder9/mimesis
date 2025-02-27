package ordermanager

import (
	"sync"
)

// OrderManager는 카페 주문을 관리하는 싱글톤 객체
// 이 객체는 애플리케이션 전역에서 단 하나만 생성되어야 함
type OrderManager struct {
	orders []string // 주문 목록을 저장하는 슬라이스
}

// OrderManager의 싱글톤 인스턴스를 저장함
var instance *OrderManager

// sync.Once를 사용해 인스턴스가 단 한번만 초기화되도록 보장함
var once sync.Once

// OrderManager의 싱글톤 인스턴스를 반환함
// 최초 호출시에만 인스턴스가 생성되고 이후에는 동일한 인스턴스를 반환함
func GetInstance() *OrderManager {
	once.Do(func() {
		instance = &OrderManager{
			orders: make([]string, 0),
		}
	})
	return instance
}

// 새로운 주문을 추가함
// 인자로 받은 주문 문자열을 orders 슬라이스에 추가함
func (om *OrderManager) AddOrder(order string) {
	om.orders = append(om.orders, order)
}

// 현재까지 저장된 모든 주문 목록을 반환
func (om *OrderManager) ShowOrders() []string {
	return om.orders
}
