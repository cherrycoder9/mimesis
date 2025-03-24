package factory

import "abstract-factory/product"

// 여성 의류 제품을 생성하는 구체적 팩토리
type WomenFactory struct{}

// 새로운 여성용 팩토리 인스턴스를 생성해 ClothingFactory 인터페이스로 반환
func NewWomenFactory() ClothingFactory {
	return &WomenFactory{}
}

// 여성용 상의를 생성해 반환하는 메서드
func (w *WomenFactory) CreateTop() product.Top {
	// product 패키지의 WomenTop 구조체 인스턴스 생성
	return &product.WomenTop{}
}

// 여성용 하의를 생성해 반환하는 메서드
func (w *WomenFactory) CreateBottom() product.Bottom {
	// product 패키지의 WomenBottom 구조체 인스턴스 생성
	return &product.WomenBottom{}
}
