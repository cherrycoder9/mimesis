package composite

import "fmt"

// Leaf 구조체는 단일 메뉴 항목을 표현함
type Leaf struct {
	name  string
	price int
}

// 주어진 이름과 가격으로 Leaf 객체를 생성
func NewLeaf(name string, price int) *Leaf {
	return &Leaf{
		name:  name,
		price: price,
	}
}

// Leaf 객체의 정보를 들여쓰기를 적용해 출력하는 메서드
// 가격 정보와 함께 메뉴 항목의 이름을 출력함
func (l *Leaf) Print(indent string) {
	fmt.Printf("%s└─%s : %d원\n", indent, l.name, l.price)
}
