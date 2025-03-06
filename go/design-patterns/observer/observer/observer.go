package observer

// 모든 옵저버들이 구현해야 하는 Update 메소드를 정의함
type Observer interface {
	// Subject의 상태 변화에 따른 알림을 처리하는 역할
	Update(message string)
}
