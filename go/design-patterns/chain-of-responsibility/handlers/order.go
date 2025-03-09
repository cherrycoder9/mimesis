package handlers

import (
	"chain-of-responsibility/common"
	"fmt"
)

// 주문 처리를 담당
type OrderHandler struct {
	next common.Handler
}

func (h *OrderHandler) SetNext(handler common.Handler) common.Handler {
	h.next = handler
	return handler
}

func (h *OrderHandler) Handle(req *common.Request) {
	if req.Order == "" {
		fmt.Println("주문이 없습니다. 주문을 확인해주세요.")
		return
	}

	fmt.Printf("✅ 주문 확인 완료 [%s]\n", req.Order)

	if h.next != nil {
		h.next.Handle(req)
	}
}
