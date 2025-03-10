package util

import (
	"bufio"
	"os"
	"strings"
)

// 사용자 입력을 처리하기 위한 유틸리티 함수

// 표준 입력으로부터 사용자 입력을 받아 반환하는 함수
// prompt 파라미터를 통해 사용자에게 입력 메시지를 출력함
func ReadInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	// 한글로 된 프롬프트 메시지 출력
	println(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
