package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"singleton/ordermanager"
)

// go run main.go
func main() {
	reader := bufio.NewReader(os.Stdin)

	// OrderManager의 싱글턴 인스턴스를 가져옴
	orderMgr := ordermanager.GetInstance()

	fmt.Println("메뉴를 주문해주세요. 종료하시려면 '종료'를 입력하세요.")

	// 사용자 입력을 받으며 주문 처리
	for {
		fmt.Print("주문할 메뉴: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("입력 오류 발생:", err)
			continue
		}
		// 개행 문자 제거
		input = strings.TrimSpace(input)

		// '종료' 입력시 반복문 종료
		if input == "종료" {
			break
		}

		// 입력된 주문을 OrderManager에 추가함
		orderMgr.AddOrder(input)
		fmt.Println("주문이 접수되었습니다:", input)
	}

	fmt.Println("\n모든 주문 목록:")
	orders := orderMgr.ShowOrders()
	for i, order := range orders {
		fmt.Printf("%d. %s\n", i+1, order)
	}

	fmt.Println("주문해 주셔서 감사합니다!")
}
