package coffee

import "fmt"

// Coffee 구조체는 커피 주문 정보를 나타냄
// 이 구조체는 빌더 패턴을 통해 생성됨
type Coffee struct {
	Size        string // 커피 사이즈 (소, 중, 대)
	CoffeeType  string // 커피 종류 (에스프레소, 아메리카노 등)
	Milk        string // 우유 종류 (전지, 저지방, 두유 등)
	Shots       int    // 에스프레소 샷 수
	Syrup       string // 시럽 종류 (바닐라, 카라멜 등)
	Temperature string // 온도 (따뜻한, 아이스)
}

// String 메서드는 커피 주문 정보를 문자열로 반환함
// 최종 객체 상태를 확인하는 용도
func (c *Coffee) String() string {
	return "사이즈: " + c.Size +
		", \n종류: " + c.CoffeeType +
		", \n우유: " + c.Milk +
		", \n에스프레소 샷: " + fmt.Sprintf("%d", c.Shots) +
		", \n시럽: " + c.Syrup +
		", \n온도: " + c.Temperature
}
