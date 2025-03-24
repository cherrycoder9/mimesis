package factory

import "abstract-factory/product"

// 의류 제품군을 생성하는 추상 팩토리
// 클라이언트는 구체적인 팩토리(남,녀)에 의존하지 않고 제품을 생성할 수 있음
type ClothingFactory interface {
	// 상의 제품을 생성하는 메서드
	CreateTop() product.Top
	// 하의 제품을 생성하는 메서드
	CreateBottom() product.Bottom
}
