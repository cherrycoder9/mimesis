package menu

// 음료 메뉴를 나타냄
type DrinkMenu struct {
	items []*MenuItem
}

// 음료 메뉴를 초기화함
func NewDrinkMenu() *DrinkMenu {
	return &DrinkMenu{
		items: []*MenuItem{
			{"아메리카노", "진한 에스프레소와 뜨거운 물의 조합", 4100},
			{"카페 라떼", "에스프레소와 부드러운 스팀우유의 조화", 4600},
			{"카라멜 마키아토", "달콤한 카라멜과 에스프레소, 우유의 완벽한 만남", 5600},
			{"자바 칩 프라푸치노", "초콜릿칩과 커피, 얼음이 어우러진 블렌디드 음료", 6300},
		},
	}
}

// 음료 메뉴를 순회할 수 있는 이터레이터를 반환
func (d *DrinkMenu) CreateIterator() Iterator {
	return &DrinkMenuIterator{
		items: d.items,
		index: 0,
	}
}
