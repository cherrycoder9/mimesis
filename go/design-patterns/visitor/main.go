package main

import "fmt"

// # 방문자 패턴 사용 목적
// 데이터 구조(객체 구조)를 변경하지 않고 새로운 연산(기능)을 추가할때 유용함
// 서로 다른 클래스에 산재한 유사한 기능을 한 곳으로 모아서 관리할 수 있음
// 새로운 방문자 클래스를 추가해 기능을 쉽게 확장할 수 있음

// 메뉴들 총 가격을 계산
type PriceCalculator struct {
	Total float64
}

func (pc *PriceCalculator) VisitBeverage(b *Beverage) {
	pc.Total += b.Price
	fmt.Printf("[음료] %s의 가격은 %.2f원입니다.\n", b.Name, b.Price)
}

func (pc *PriceCalculator) VisitFood(f *Food) {
	pc.Total += f.Price
	fmt.Printf("[음식] %s의 가격은 %.2f원입니다.\n", f.Name, f.Price)
}

// 메뉴들의 영양 정보를 출력
type NutritionInfoPrinter struct{}

func (nip *NutritionInfoPrinter) VisitBeverage(b *Beverage) {
	fmt.Printf("[음료] %s의 칼로리는 %dkcal입니다. (시즌 한정: %v)\n",
		b.Name, b.Calories, b.IsSeasonal)
}

func (nip *NutritionInfoPrinter) VisitFood(f *Food) {
	fmt.Printf("[음식] %s의 칼로리는 %dkcal입니다. (비건 메뉴: %v)\n",
		f.Name, f.Calories, f.isVegan)
}

// go run .
func main() {
	// 메뉴 아이템 생성
	items := []MenuItem{
		&Beverage{"아메리카노", 4500, 10, false},
		&Beverage{"딸기 라떼", 6500, 250, true},
		&Food{"클럽 샌드위치", 7500, 500, false},
		&Food{"비건 샐러드", 6800, 320, true},
	}

	// 가격 계산 방문자 생성 및 적용
	priceCalculator := &PriceCalculator{}
	for _, item := range items {
		item.Accept(priceCalculator)
	}
	fmt.Printf("\n 총 결제 금액: %.2f원\n\n", priceCalculator.Total)

	// 영양 정보 출력 방문자 생성 및 적용
	NutritionInfoPrinter := &NutritionInfoPrinter{}
	for _, item := range items {
		item.Accept(NutritionInfoPrinter)
	}
}
