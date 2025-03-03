package beverage

// 기본 음료인 에스프레소를 나타냄
type Espresso struct{}

// NewEspresso는 새로운 에스프레소 객체를 생성해 반환함
func NewEspresso() Beverage {
	return &Espresso{}
}

// GetDescription은 에스프레소의 설명을 반환함
func (e *Espresso) GetDescription() string {
	return "에스프레소"
}

// 에스프레소의 가격 반환
func (e *Espresso) Cost() float64 {
	return 1500.0
}
