package main

import (
	"abstract-factory/factory"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// # 추상팩토리 패턴
// 관련있는 객체 집합을 생성하기 위한 인터페이스를 제공
// 구체적인 클래스를 지정하지 않음
// 서로 다른 제품군을 일관되게 생성할때 효과적
func main() {
	fmt.Println("나인원 의류 쇼핑몰에 오신 것을 환영합니다.")
	fmt.Println("원하시는 컬렉션을 선택하세요. (남성 / 여성):")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var clothFactory factory.ClothingFactory

	// 입력에 따른 팩토리 인스턴스 생성 (남성용, 여성용)
	if strings.EqualFold(input, "남성") {
		clothFactory = factory.NewMenFactory()
	} else if strings.EqualFold(input, "여성") {
		clothFactory = factory.NewWomenFactory()
	} else {
		fmt.Println("잘못된 입력입니다. 프로그램을 종료합니다.")
		return
	}

	// 팩토리를 통해 상의와 하의 제품 생성
	top := clothFactory.CreateTop()
	bottom := clothFactory.CreateBottom()

	// 생성된 제품들의 상태를 출력
	fmt.Println("생성된 제품들:")
	fmt.Println(top.Wear())
	fmt.Println(bottom.Wear())
}
