// 맵 디스패처 패턴
package main

import "fmt"

// 연산 함수들을 위한 별칭 타입
type calculateFunc func(int, int) int

// 기본 연산 함수들 정의
func add(a, b int) int  { return a + b }
func sub(a, b int) int  { return a - b }
func mult(a, b int) int { return a * b }
func div(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

// 문자열 키를 연산 함수에 바인딩한 맵
var operations = map[string]calculateFunc{
	"+":  add,
	"-":  sub,
	"*":  mult,
	"/":  div,
	"<<": func(a, b int) int { return a << b },
	">>": func(a, b int) int { return a >> b },
}

// 계산 함수
func calculateWithMap(a, b int, opString string) int {
	if operation, ok := operations[opString]; ok {
		return operation(a, b)
	}
	panic("operation not supported")
}

func main() {
	fmt.Println(calculateWithMap(3, 1, "+"))
}
