package implementation

// 커피 추출 방식을 정의하는 인터페이스
// 다양한 커피 추출 방법을 독립적으로 추가할 수 있음
type BrewMethod interface {
	Brew() string
}

// 에스프레소 머신을 통한 커피 추출 방식을 구현하는 구조체
type EspressoMachine struct{}

// 새로운 EspressoMachine 인스턴스를 생성하는 함수
// Go 언어에선 객체 지향 언어에서의 생성자 개념이 없기 때문에,
// 생성 역할을 하는 별도의 함수를 사용하는 것이 관례
func NewEspressoMachine() *EspressoMachine {
	return &EspressoMachine{}
}

// 에스프레소 머신을 사용한 커피 추출 방식을 반환하는 메서드
func (e *EspressoMachine) Brew() string {
	return "에스프레소 머신을 사용해 커피를 추출했습니다."
}

// 핸드드립 방식을 통한 커피 추출 방식 구현하는 구조체
type HandBrew struct{}

// 새로운 HandBrew 인스턴스를 생성하는 함수
func NewHandBrew() *HandBrew {
	return &HandBrew{}
}

// 핸드드립 방식을 사용한 커피 추출 방식을 반환
func (h *HandBrew) Brew() string {
	return "핸드드립 방식을 사용해 커피를 추출했습니다."
}
