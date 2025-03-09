package handlers

import (
	"chain-of-responsibility/common"
	"fmt"
)

// 결제를 처리하는 구조체
type PaymentHandler struct {
	next common.Handler
}

func (h *PaymentHandler) SetNext(handler common.Handler) common.Handler {
	h.next = handler
	return handler
}

func (h *PaymentHandler) Handle(req *common.Request) {
	if !req.Paid {
		fmt.Println("❌ 결제가 완료되지 않았습니다. 결제를 먼저 진행해주세요.")
		return
	}

	fmt.Println("💳 결제 확인 완료")

	if h.next != nil {
		h.next.Handle(req)
	}
}
