package handlers

import (
	"chain-of-responsibility/common"
	"fmt"
)

// 멤버십 포인트 적립을 처리하는 구조체
type RewardHandler struct {
	next common.Handler
}

func (h *RewardHandler) SetNext(handler common.Handler) common.Handler {
	h.next = handler
	return handler
}

func (h *RewardHandler) Handle(req *common.Request) {
	if req.MemberID != "" {
		fmt.Printf("🎉 멤버십 포인트가 적립되었습니다. (회원 ID: %s)\n", req.MemberID)
	} else {
		fmt.Println("ℹ️ 비회원입니다. 멤버십 포인트는 적립되지 않습니다.")
	}

	fmt.Println("☕️ 음료 준비 완료! 맛있게 드세요!")
}
