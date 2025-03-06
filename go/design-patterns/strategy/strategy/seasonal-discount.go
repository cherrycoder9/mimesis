package strategy

// 계절 할인 전략
type SeasonalDiscountStrategy struct {
	DiscountRate float64 // 할인율 (0.1은 10% 할인)
}

// 할인율을 적용해 최종 가격 반환
func (s *SeasonalDiscountStrategy) Calculate(price float64) float64 {
	return price * (1 - s.DiscountRate)
}
