package main

import "fmt"

// Go에서는 기본 타입에 메서드를 직접 추가할 수 없다
// 그러나 type alias(타입 별칭)를 사용하면 가능해진다 ?? 타입 별칭은 type age = uint 임
// 타입 정의가 type age uint
type age uint
type newAge = uint

func (a age) valid() bool {
	return a < 120
}

func main() {
	// age 타입이 객체처럼 동작함
	// 원시 타입에는 직접 메서드 정의 불가
	// type alias는 Go에서 원시 타입을 확장하는 방법 (사실 타입 정의임, 저자가 실수한듯)
	// 타입 별칭과 타입 정의는 다름
	var myAge uint = 20
	var hisAge age = 20
	var herAge newAge = 20

	// 타입 정의인 경우 같은 타입이 아니므로 형변환을 해야함
	fmt.Println(myAge == uint(hisAge))
	// 타입 별칭인 경우 형변환 필요없이 동일타입 취급됨
	fmt.Println(myAge == herAge)
	// 타입 정의인 경우 새 타입을 만드는 것이므로 메서드를 붙일수도 있음
	fmt.Println(hisAge.valid())
}
