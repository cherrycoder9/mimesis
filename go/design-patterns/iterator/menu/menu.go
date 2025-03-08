package menu

// 메뉴 항목들을 순회하는 메서드 정의
type Iterator interface {
	HasNext() bool
	Next() *MenuItem
}

// 메뉴에 이터레이터를 반환하는 메서드를 정의
type Menu interface {
	CreateIterator() Iterator
}

// MenuItem은 메뉴의 한 항목을 나타냄
type MenuItem struct {
	Name        string
	Description string
	Price       float64
}
