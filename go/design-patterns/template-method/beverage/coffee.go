package beverage

import "fmt"

// Beverage 인터페이스를 구현하는 커피 음료
type Coffee struct{}

// 커피 원두를 우려내는 단계 구현
func (c *Coffee) Brew() {
	fmt.Println("   커피 원두를 우려냅니다.")
}

// 커피에 설탕과 크림을 추가하는 단계 구현
func (c *Coffee) AddCondiments() {
	fmt.Println("   설탕과 크림을 추가합니다.")
}

// 사용자 입력에 따라 첨가물 추가 여부를 결정함
func (c *Coffee) CustomerWantsCondiments() bool {
	var answer string
	fmt.Print("   첨가물을 추가하시겠습니까? (예/아니오): ")
	fmt.Scanln(&answer)
	return answer == "예" || answer == "네"
}
