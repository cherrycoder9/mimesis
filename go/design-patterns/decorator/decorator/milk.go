package decorator

import "decorator/beverage"

// 음료에 우유를 추가하는 데코레이터
// 기존 음료(Beverage 인터페이스)를 포함해 추가 기능 제공
type MilkDecorator struct {
	Beverage beverage.Beverage
}

// 기존 음료에 우유 데코레이터를 적용한 새로운 음료 객체를 생성
func NewMilkDecorator(b beverage.Beverage) beverage.Beverage {
	return &MilkDecorator{Beverage: b}
}

// 기존 음료의 설명에 우유를 추가해 반환
func (m *MilkDecorator) GetDescription() string {
	return m.Beverage.GetDescription() + " + 우유"
}

// 기존 음료 가격에 우유 추가 비용을 더해 반환
func (m *MilkDecorator) Cost() float64 {
	return m.Beverage.Cost() + 500.0
}
