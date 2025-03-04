package beverage

import "fmt"

// Beverage 인터페이스를 구현하는 차 음료
type Tea struct{}

// 찻잎을 우려내는 단계 구현
func (t *Tea) Brew() {
	fmt.Println("   찻잎을 우려냅니다.")
}

// 차에 레몬을 추가하는 단계 구현
func (t *Tea) AddCondiments() {
	fmt.Println("   레몬을 추가합니다.")
}

// 사용자 입력에 따라 첨가물 추가 여부를 결정함
func (t *Tea) CustomerWantsCondiments() bool {
	var answer string
	fmt.Print("   첨가물을 추가하시겠습니까? (예/아니오): ")
	fmt.Scanln(&answer)
	return answer == "예" || answer == "네"
}
