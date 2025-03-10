package composite

import "fmt"

// 자식 메뉴 항목들을 가지며, MenuComponent 인터페이스를 구현하는 구조체
// 복합 메뉴 항목은 내부에 여러 개의 하위 메뉴 항목을 포함할 수 있음
type Composite struct {
	name     string
	children []MenuComponent
}

// 주어진 이름으로 Composite 객체를 생성하는 함수
func NewComposite(name string) *Composite {
	return &Composite{
		name:     name,
		children: make([]MenuComponent, 0),
	}
}

// Composite 객체에 새로운 자식 메뉴 항목을 추가하는 메서드
func (c *Composite) Add(component MenuComponent) {
	c.children = append(c.children, component)
}

// Composite 객체와 그 하위 항목들을 들여쓰기를 적용해 출력하는 메서드
// 전체 메뉴 구조를 트리 형태로 쉽게 확인할 수 있음
func (c *Composite) Print(indent string) {
	fmt.Println(indent + "└─" + c.name)
	for _, child := range c.children {
		child.Print(indent + "    ")
	}
}
