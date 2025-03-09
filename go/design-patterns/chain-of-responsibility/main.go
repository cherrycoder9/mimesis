package main

import (
	"chain-of-responsibility/common"
	"chain-of-responsibility/handlers"
	"fmt"
)

// # 책임 연쇄 패턴
// 요청을 처리할 수 있는 객체들을 사슬로 연결하고 요청을 처리할 수 있을 때까지 이 객체들 간에 요청을 전달하는 행동 디자인 패턴. 체인이 길어지면 디버깅이 어려워지고 성능이 저하될 수 있다. 요청이 항상 특정 객체에서 처리되는 경우에는 불필요하다.
func main() {
	// 핸들러 초기화 및 연결
	orderHandler := &handlers.OrderHandler{}
	paymentHandler := &handlers.PaymentHandler{}
	rewardHandler := &handlers.RewardHandler{}

	orderHandler.SetNext(paymentHandler).SetNext(rewardHandler)

	// 고객 요청 예시
	requests := []*common.Request{
		{Order: "아메리카노", Paid: true, MemberID: "STB123"},
		{Order: "라떼", Paid: false, MemberID: "STB999"},
		{Order: "", Paid: true, MemberID: ""},
	}

	fmt.Printf("🏪 카페베네 매장 주문 처리 시스템\n\n")

	for i, req := range requests {
		fmt.Printf("----- [%d번째 고객] -----\n", i+1)
		orderHandler.Handle(req)
		fmt.Println()
	}
}
