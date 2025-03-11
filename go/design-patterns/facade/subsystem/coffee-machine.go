package subsystem

import "fmt"

// 커피를 추출하는 기능을 담당하는 구조체
type CoffeeMachine struct{}

// CoffeeMachine 객체를 생성해 반환함
func NewCoffeeMachine() *CoffeeMachine {
	return &CoffeeMachine{}
}

// 커피를 추출하는 기능 수행
func (cm *CoffeeMachine) Brew() {
	fmt.Println("커피머신, 커피 추출 중...")
}
