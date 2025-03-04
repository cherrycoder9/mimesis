package strategy

// 할인 계산 전략을 위한 인터페이스, 여러 할인 정책을 캡슐화함
type DiscountStrategy interface {
	// 원래 가격을 받아 할인된 가격 반환
	Calculate(price float64) float64
}
