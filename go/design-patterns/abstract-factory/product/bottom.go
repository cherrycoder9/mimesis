package product

// 하의 제품에 대한 공통 기능 정의
type Bottom interface {
	// 제품 착용시 상태 메시지 반환
	Wear() string
}

// 남성용 하의 제품을 나타내는 구조체
type MenBottom struct{}

// 남성 하의를 착용했을 때의 메시지를 반환하는 메서드
func (mb *MenBottom) Wear() string {
	return "남성 하의를 착용했습니다."
}

// 여성용 하의 제품을 나타내는 구조체
type WomenBottom struct{}

// 여성 하의를 착용했을 때의 메시지를 반환하는 메서드
func (wb *WomenBottom) Wear() string {
	return "여성 하의를 착용했습니다."
}
