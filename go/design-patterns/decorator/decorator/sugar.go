package decorator

import "decorator/beverage"

// 음료에 설탕을 추가하는 데코레이터
// 기존 음료 인터페이스를 포함해 설탕 추가 기능 제공
type SugarDecorator struct {
	Beverage beverage.Beverage
}

// 기존 음료에 설탕 데코레이터를 적용한 새로운 음료 객체 생성
func NewSugarDecorator(b beverage.Beverage) beverage.Beverage {
	return &SugarDecorator{Beverage: b}
}

// 기존 음료 설명에 설탕을 추가해 반환
func (s *SugarDecorator) GetDescription() string {
	return s.Beverage.GetDescription() + " + 설탕"
}

func (s *SugarDecorator) Cost() float64 {
	return s.Beverage.Cost() + 300.0
}
