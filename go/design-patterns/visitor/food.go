package main

// 음식 메뉴를 나타냄
type Food struct {
	Name     string
	Price    float64
	Calories int
	isVegan  bool // 비건 여부
}

// 방문자를 받아 해당 방문자의 메서드 호출
func (f *Food) Accept(v Visitor) {
	v.VisitFood(f)
}
