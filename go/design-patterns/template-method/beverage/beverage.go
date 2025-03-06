package beverage

import "fmt"

// Beverage 인터페이스는 음료 제조에 필요한 단계(메서드)를 정의함
// 각 구체적 음료는 이 인터페이스를 구현해 자신만의 방식으로 처리
type Beverage interface {
	// 재료를 우려내는 단계 (각 음료마다 다르게 구현)
	Brew()
	// 첨가물을 추가하는 단계 (각 음료마다 다르게 구현)
	AddCondiments()
	// 후크 메서드: 사용자가 첨가물을 원하는지 여부를 결정
	// 기본 템플릿 메서드에서 이 값을 통해 첨가물 단계를 수행할지 결정
	CustomerWantsCondiments() bool
}

// 템플릿 메서드, 음료 제조의 전체 순서를 정의함
// 세부 단계는 Beverage 인터페이스를 구현한 객체에 위임함
func PrepareRecipe(b Beverage) {
	fmt.Println("1. 물을 끓입니다.")
	boilWater()
	fmt.Println("2. 재료를 우려냅니다.")
	b.Brew()
	fmt.Println("3. 컵에 붓습니다.")
	pourInCup()
	// 후크 메서드의 결과에 따라 첨가물 추가 여부 결정
	if b.CustomerWantsCondiments() {
		fmt.Println("4. 첨가물을 추가합니다.")
		b.AddCondiments()
	} else {
		fmt.Println("4. 첨가물 없이 제공합니다.")
	}
}

// 물을 끓이는 공통 단계를 구현
func boilWater() {
	fmt.Println("   물 끓이는 중... (대기)")
}

// 음료를 컵에 붓는 공통 단계를 구현
func pourInCup() {
	fmt.Println("   컵에 붓는 중... (대기)")
}
