package main

import "composite/composite"

// # 컴포짓 패턴
// 객체를 트리 구조로 구성해 개별 객체와 집합 객체를 동일하게 다룸
// 파일 시스템 등에서 파일과 폴더를 일관되게 처리할 수 있음
// 계층적인 구조를 관리할 때 유용함
// 하지만 모든 객체가 동일 인터페이스를 강제받아 설계가 제한될 수 있음
// 계층 구조가 없거나 단순한 리스트로 충분한 경우에는 필요없음
func main() {
	// 단일 메뉴 항목 생성 (Leaf)
	espresso := composite.NewLeaf("에스프레소", 2500)
	americano := composite.NewLeaf("아메리카노", 3000)
	latte := composite.NewLeaf("라떼", 3500)

	// 복합 메뉴 항목 생성 (Composite) - 커피 메뉴
	coffeeMenu := composite.NewComposite("커피 메뉴")
	coffeeMenu.Add(espresso)
	coffeeMenu.Add(americano)
	coffeeMenu.Add(latte)

	// 또 다른 단일 메뉴 항목 생성
	greenTea := composite.NewLeaf("녹차", 2800)
	herbalTea := composite.NewLeaf("허브티", 3200)

	// 복합 메뉴 항목 생성 (Composite) - 차 메뉴
	teaMenu := composite.NewComposite("차 메뉴")
	teaMenu.Add(greenTea)
	teaMenu.Add(herbalTea)

	// 전체 메뉴 구성 (Composite)
	fullMenu := composite.NewComposite("전체 메뉴")
	fullMenu.Add(coffeeMenu)
	fullMenu.Add(teaMenu)

	// 메뉴 출력
	println("카페 메뉴 출력:")
	// 들여쓰기를 통해 트리 형태로 출력
	fullMenu.Print("  ")
}
