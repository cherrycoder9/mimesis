package main

import (
	"bufio"
	"fmt"
	"os"
	"proxy/cafe"
	"strings"
)

// # 프록시 패턴
// 다른 객체에 대한 접근을 제어하거나 대리 역할을 수행하는 객체를 제공
// 실제 객체의 초기화 비용이 크거나 보안, 캐싱이 필요할 때 유용함
// 프록시는 클라이언트와 실제 객체 사이에서 중간자 역할을 하며 요청을 가로챔
func main() {
	// 프록시 객체 생성 (로그인 상태에 따라 authorized를 변경 가능)
	proxy := cafe.NewProxy(authenticate())

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("주문할 음료를 입력해주세요: ")
	if scanner.Scan() {
		drink := scanner.Text()
		proxy.Order(drink)
	}
}

// 간단한 로그인 인증 함수
func authenticate() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("로그인이 필요합니다. 사용자 이름을 입력하세요 (빈칸 입력시 비회원): ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())
	if username == "" {
		fmt.Println("비회원으로 접속하셨습니다.")
		return false
	}
	fmt.Printf("'%s' 님 로그인 성공!\n", username)
	return true
}
