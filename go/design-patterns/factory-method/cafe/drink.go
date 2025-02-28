package cafe

import "fmt"

// Drink 인터페이스는 모든 음료가 구현해야 할 공통 동작 정의
// 여기선 음료를 준비하는 동작을 Prepare() 메서드로 정의함
type Drink interface {
	Prepare()
}

// Coffee 구조체는 커피 음료를 나타내며 Drink 인터페이스를 구현함
type Coffee struct{}

// Prepare 메서드는 커피 준비 과정을 출력함
func (c *Coffee) Prepare() {
	fmt.Println("커피 준비 중...")
}

// Tea 구조체는 차 음료를 나타내며 Drink 인터페이스 구현
type Tea struct{}

// Prepare 메서드는 차 주비 과정을 출력
func (t *Tea) Prepare() {
	fmt.Println("차 준비 중...")
}
