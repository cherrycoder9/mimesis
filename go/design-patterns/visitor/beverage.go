package main

// 음료 메뉴를 나타냄
type Beverage struct {
	Name       string
	Price      float64
	Calories   int
	IsSeasonal bool // 시즌 한정 메뉴 여부
}

// 방문자를 받아 해당 방문자의 메서드 호출
func (b *Beverage) Accept(v Visitor) {
	v.VisitBeverage(b)
}
