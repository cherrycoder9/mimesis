package beverage

// 음료의 기본 메서드를 정의하는 인터페이스
// 음료의 설명과 가격 정보를 가져올 수 있음
type Beverage interface {
	// 음료에 대한 설명을 반환
	GetDescription() string
	// 음료의 가격 반환
	Cost() float64
}
