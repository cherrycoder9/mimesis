package facade

import (
	"facade/subsystem"
	"fmt"
)

// 커피 주문을 위한 서브시스템들을 단순화하는 역할을 함
type CoffeeFacade struct {
	coffeeMachine  *subsystem.CoffeeMachine  // 커피 추출 기능
	milkStreamer   *subsystem.MilkSteamer    // 우유 스팀 기능
	sugarDispenser *subsystem.SugarDispenser // 설탕 공급 기능
}

// CoffeeFacade 객체를 초기화해 반환함
func NewCoffeeFacade() *CoffeeFacade {
	return &CoffeeFacade{
		coffeeMachine:  subsystem.NewCoffeeMachine(),
		milkStreamer:   subsystem.NewMilkSteamer(),
		sugarDispenser: subsystem.NewSugarDispenser(),
	}
}

// 아메리카노 주문 과정을 단순화해 실행하는 메서드
func (f *CoffeeFacade) OrderAmericano() {
	// 원두를 갈고 커피 추출
	fmt.Println("커피머신, 원두를 갈고 추출합니다.")
	f.coffeeMachine.Brew()

	// 우유를 데우고 거품을 생성
	fmt.Println("우유 스티머, 우유를 데우고 거품을 만듭니다.")
	f.milkStreamer.Steam()

	// 설탕 추가
	fmt.Println("설탕 디스펜서, 설탕을 추가합니다.")
	f.sugarDispenser.Dispense()

	fmt.Println("아메리카노 준비 완료!!")

}
