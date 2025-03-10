package interpreter

import "unicode"

// 입력 문자열을 토큰으로 분해하는 렉서(Tokenizer) 및 토큰 정의

// TokenType, 토큰의 종류를 문자열 상수로 정의함
type TokenType string

// string 으로 안하고 굳이 TokenType을 만든 이유?
// 사용자 정의 타입을 사용하면 함수나 메서드의 매개변수에 해당 타입을
// 요구할 수 있어 실수로 다른 문자열 값을 전달하는 것을 방지할 수 있음
// 해당 값들이 토큰의 종류를 나타낸다는 의미 전달을 명확하게 함
// 이후에 토큰 타입에 관련된 메서드를 추가해야 하는 경우, 사용자 정의 타입에
// 메서드를 붙여서 처리할 수 있음.
const (
	NUMBER TokenType = "NUMBER" // 숫자
	PLUS   TokenType = "PLUS"   // +
	MINUS  TokenType = "MINUS"  // -
	MUL    TokenType = "MUL"    // *
	DIV    TokenType = "DIV"    // /
	LPAREN TokenType = "LPAREN" // (
	RPAREN TokenType = "RPAREN" // )
	EOF    TokenType = "EOF"    // 입력 종료 표시
)

// Token 구조체는 토큰의 타입과 값을 포함
type Token struct {
	Type  TokenType
	Value string
}

// Lex 함수는 입력 문자열을 읽어 토큰의 슬라이스로 분해함
// 이 함수는 간단한 렉서(Tokenizer) 역할을 수행함
func Lex(input string) []Token {
	tokens := []Token{}
	i := 0
	for i < len(input) {
		c := input[i]
		// 공백은 무시
		if c == ' ' {
			i++
			continue
		}
		// 숫자 처리, 연속된 숫자들을 하나의 토큰으로 인식
		if unicode.IsDigit(rune(c)) {
			numStr := ""
			for i < len(input) && unicode.IsDigit(rune(input[i])) {
				numStr += string(input[i])
				i++
			}
			tokens = append(tokens, Token{Type: NUMBER, Value: numStr})
			continue
		}
		// 연산자 및 괄호 처리
		switch c {
		case '+':
			tokens = append(tokens, Token{Type: PLUS, Value: string(c)})
		case '-':
			tokens = append(tokens, Token{Type: MINUS, Value: string(c)})
		case '*':
			tokens = append(tokens, Token{Type: MUL, Value: string(c)})
		case '/':
			tokens = append(tokens, Token{Type: DIV, Value: string(c)})
		case '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: string(c)})
		case ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: string(c)})
		default:
			// pass
		}
		i++
	}
	// 입력 종료 토큰 추가
	tokens = append(tokens, Token{Type: EOF, Value: ""})
	return tokens
}
