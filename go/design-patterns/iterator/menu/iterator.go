package menu

// 음료 메뉴를 순회하는 이터레이터 구현체
type DrinkMenuIterator struct {
	items []*MenuItem
	index int
}

// 다음 메뉴 항목이 존재하는지 확인
func (it *DrinkMenuIterator) HasNext() bool {
	return it.index < len(it.items)
}

// 다음 메뉴 항목을 반환
func (it *DrinkMenuIterator) Next() *MenuItem {
	if it.HasNext() {
		item := it.items[it.index]
		it.index++
		return item
	}
	return nil
}
