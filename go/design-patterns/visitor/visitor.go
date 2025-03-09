package main

// 다양한 메뉴 항목을 처리할 메서드를 정의함
type Visitor interface {
	VisitBeverage(b *Beverage)
	VisitFood(f *Food)
}
