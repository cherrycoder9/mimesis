package common

// 고객의 요청을 표현하는 구조체
type Request struct {
	Order    string // 주문한 음료
	Paid     bool   // 결제 완료 여부
	MemberID string // 멤버십 ID (없으면 빈 문자열)
}

// 요청을 처리하는 인터페이스
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(req *Request)
}
