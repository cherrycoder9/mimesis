package main

import (
	"fmt"
	"interpreter/interpreter"
	"interpreter/util"
)

// # 인터프리터 패턴
// 특정 문법이나 언어를 해석하고 실행할 수 있는 방법 정의
// 간단한 언어 규칙을 해석해야 할때 유용함
// 쿼리 언어나 수식 계산기를 만들때 적용할 수 있음
// 각 문법 요소를 객체로 표현하며 그것들을 조합해 해석함
func main() {
	fmt.Println("==== 카페베네 계산기 ====")

	// 사용자로부터 수식을 입력 받음
	input := util.ReadInput("계산할 수식을 입력하세요 (예: 3 + 5 * (3 - 1)): ")

	// 입력받은 문자열을 렉서를 통해 토큰으로 분해
	tokens := interpreter.Lex(input)
	// 파서를 생성하고 토큰을 기반으로 구문 분석(파싱)하고 표현식 트리(AST) 생성
	parser := interpreter.NewParser(tokens)
	ast := parser.Parse()

	// 생성된 AST를 해석(Interpret)해 최종 결과를 도출
	result := ast.Interpret()

	// 결과 출력
	fmt.Println("계산 결과:", result)
}
