package subject

import "observer/observer"

// 커피 주문 상태를 관리하고, 상태 변화가 있을때 등록된 옵저버들에게 알림을 보냄
type OrderSubject struct {
	observers []observer.Observer // 등록된 옵저버 목록
	order     string              // 현재 주문 상태 (예: 주문한 커피 종류)
}

// OrderSubject 객체를 생성하는 생성자
func NewOrderSubject() *OrderSubject {
	return &OrderSubject{
		observers: make([]observer.Observer, 0),
	}
}

// 새로운 옵저버를 등록
func (os *OrderSubject) Registers(obs observer.Observer) {
	os.observers = append(os.observers, obs)
}

// 등록된 옵저버를 제거함
func (os *OrderSubject) Deregister(obs observer.Observer) {
	for i, o := range os.observers {
		if o == obs {
			// 슬라이스에서 해당 옵저버 제거
			os.observers = append(os.observers[:i], os.observers[i+1:]...)
			break
		}
	}
}

// 상태 변화가 있을때 모든 옵저버에게 알림을 보냄
func (os *OrderSubject) NotifyObservers(message string) {
	for _, obs := range os.observers {
		obs.Update(message)
	}
}

// 새로운 주문을 받았을때 호출됨, 주문 상태 업데이트하고 옵저버들에게 알림 발송
func (os *OrderSubject) NewOrder(order string) {
	os.order = order
	// 모든 옵저버에게 주문 상태 변경 알림 전파
	os.NotifyObservers("새 주문 접수: " + order)
}
