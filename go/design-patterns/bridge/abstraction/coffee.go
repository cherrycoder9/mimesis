package abstraction

import "bridge/implementation"

// 커피의 추상적 개념을 나타내는 구조체
// BrewMethod 인터페이스를 통해 구체적인 커피 추출 방식을 연결함
type Coffee struct {
	Name       string
	BrewMethod implementation.BrewMethod // 커피 추출 구현체 (브리지 역할)
}

// 새로운 Coffee 인스턴스를 생성하는 함수
func NewCoffee(name string, brewMethod implementation.BrewMethod) *Coffee {
	return &Coffee{
		Name:       name,
		BrewMethod: brewMethod,
	}
}

// 커피 제공 과정 메서드
// 내부적으로 BrewMethod의 Brew()를 호출해 구체적인 추출 방법에 따른 결과를 포함
func (c *Coffee) Serve() string {
	brewMessage := c.BrewMethod.Brew() // 구체적인 추출 방법
	return "선택하신 " + c.Name + "가 준비되었습니다. " + brewMessage
}
