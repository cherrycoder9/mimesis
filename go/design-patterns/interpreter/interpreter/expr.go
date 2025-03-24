package interpreter

// Expression 인터페이스는 모든 표현식이 구현해야 하는 Interpret 메소드를 정의
// 인터프리터 패턴의 핵심으로, 각 구체적인 표현식은 이 메소드를 통해 자신의 결과를 계산
type Expression interface {
	Interpret() int
}

// NumberExpression은 숫자(터미널 표현식)를 나타냄
type NumberExpression struct {
	Value int
}

// 숫자 자체의 값을 반환하는 메서드
func (n *NumberExpression) Interpret() int {
	return n.Value
}

// AddExpression은 덧셈 연산을 나타냄
type AddExpression struct {
	Left, Right Expression // 왼쪽 오른쪽 피연산자
}

// 왼쪽과 오른쪽 표현식의 해석 결과를 더하는 메서드
func (a *AddExpression) Interpret() int {
	return a.Left.Interpret() + a.Right.Interpret()
}

// 뺄셈 연산을 나타내는 구조체
type SubtractExpression struct {
	Left, Right Expression // 왼쪽, 오른쪽 피연산자
}

// 왼쪽 표현식에서 오른쪽 표현식의 해석 결과를 빼는 메서드
func (s *SubtractExpression) Interpret() int {
	return s.Left.Interpret() - s.Right.Interpret()
}

// 곱셈 연산을 나타내는 구조체
type MultiplyExpression struct {
	Left, Right Expression // 왼쪽, 오른쪽 피연산자
}

// 왼쪽과 오른쪽 표현식의 해석 결과를 곱함
func (m *MultiplyExpression) Interpret() int {
	return m.Left.Interpret() * m.Right.Interpret()
}

// 나눗셈 연산을 나타내는 구조체
type DivideExpression struct {
	Left, Right Expression // 왼쪽 오른쪽 피연산자
}

// 왼쪽 표현식의 해석 결과를 오른쪽 표현식의 해석 결과로 나눔
func (d *DivideExpression) Interpret() int {
	return d.Left.Interpret() / d.Right.Interpret()
}
