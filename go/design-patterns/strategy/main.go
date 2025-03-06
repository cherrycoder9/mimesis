package main

import (
	"bufio"
	"fmt"
	"os"
	"strategy/strategy"
	"strconv"
	"strings"
)

// 특정 작업을 수행하는 알고리즘을 런타임에 교체, 행위를 캡슐화
// 알고리즘이 자주 바뀌거나 조건에 따라 다른 방식을 적용해야 할때 사용
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("주문 금액을 입력하세요: ")
	inputPrice, _ := reader.ReadString('\n')
	inputPrice = strings.TrimSpace(inputPrice)
	price, err := strconv.ParseFloat(inputPrice, 64)
	if err != nil {
		fmt.Println("올바른 숫자를 입력하세요.")
		return
	}

	// 할인 전략 선택 메뉴
	fmt.Println("할인 전략을 선택하세요:")
	fmt.Println("1. 무할인")
	fmt.Println("2. 계절할인")

	fmt.Print("옵션 번호를 입력하세요 (1 또는 2): ")
	optionStr, _ := reader.ReadString('\n')
	optionStr = strings.TrimSpace(optionStr)
	option, err := strconv.Atoi(optionStr)
	if err != nil {
		fmt.Println("올바른 옵션 번호를 입력하세요.")
		return
	}

	// 선택한 옵션에 따라 할인 전략 생성
	var discountStrategy strategy.DiscountStrategy

	switch option {
	case 1:
		discountStrategy = &strategy.NoDiscountStrategy{}
	case 2:
		discountStrategy = &strategy.SeasonalDiscountStrategy{DiscountRate: 0.1}
	default:
		fmt.Println("올바른 옵션 번호가 아닙니다.")
		return
	}

	// 할인 전략을 적용해 최종 가격 계산
	finalPrice := discountStrategy.Calculate(price)

	// 결과 출력
	fmt.Printf("원래 가격: %.2f원, 할인 적용 후 가격: %.2f원\n", price, finalPrice)
}
