package strategy

// 할인을 적용하지 않는 전략
type NoDiscountStrategy struct{}

// 할인 없이 원래 가격을 그대로 반환
func (n *NoDiscountStrategy) Calculate(price float64) float64 {
	return price
}
