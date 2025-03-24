package factory

import "abstract-factory/product"

// 남성 의류제품을 생성하는 구체적 팩토리
type MenFactory struct{}

// 새로운 남성용 팩토리 인스턴스를 생성해 ClothingFactory 인터페이스로 반환함
func NewMenFactory() ClothingFactory {
	return &MenFactory{}
}

// 남성용 상의를 생성해 반환하는 메서드
func (m *MenFactory) CreateTop() product.Top {
	// product 패키지의 MenTop 구조체 인스턴스 생성
	return &product.MenTop{}
}

// 남성용 하의를 생성해 반환하는 메서드
func (m *MenFactory) CreateBottom() product.Bottom {
	// product 패키지의 MenBottom 구조체 인스턴스 생성
	return &product.MenBottom{}
}
