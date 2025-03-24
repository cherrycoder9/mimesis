package product

/*
	Top 인터페이스는 상의 제품에 대한 공통 기능을 정의함
	실제 남성용, 여성용 상의 제품은 이 인터페이스를 구현함
*/
type Top interface {
	// Wear 메서드는 제품을 착용했을 때의 상태 메시지를 반환함
	Wear() string
}

// MenTop은 남성용 상의 제품을 나타내는 구조체
type MenTop struct{}

// 남성 성의를 착용했을 때의 메시지를 반환하는 메서드
func (mt *MenTop) Wear() string {
	return "남성 상의를 착용했습니다."
}

// 여성용 상의 제품을 나타내는 구조체
type WomenTop struct{}

// 여성 상의를 착용했을때의 메시지를 반환하는 메서드
func (wt *WomenTop) Wear() string {
	return "여성 상의를 착용했습니다."
}
