package coffee

// Coffee 객체를 단계별로 생성하기 위한 빌더
type CoffeeBuilder struct {
	coffee *Coffee
}

// 새로운 CoffeeBuilder 객체를 초기화해 반환함
func NewCoffeeBuilder() *CoffeeBuilder {
	return &CoffeeBuilder{coffee: &Coffee{}}
}

// 커피의 사이즈를 설정함
func (cb *CoffeeBuilder) SetSize(size string) *CoffeeBuilder {
	cb.coffee.Size = size
	return cb
}

// 커피의 종류를 설정함
func (cb *CoffeeBuilder) SetCoffeeType(coffeeType string) *CoffeeBuilder {
	cb.coffee.CoffeeType = coffeeType
	return cb
}

// 커피에 들어갈 우유 종류를 설정
func (cb *CoffeeBuilder) SetMilk(milk string) *CoffeeBuilder {
	cb.coffee.Milk = milk
	return cb
}

// 에스프레소 샷 수를 추가
func (cb *CoffeeBuilder) AddShots(shots int) *CoffeeBuilder {
	cb.coffee.Shots += shots
	return cb
}

// 시럽 종류를 설정
func (cb *CoffeeBuilder) SetSyrup(syrup string) *CoffeeBuilder {
	cb.coffee.Syrup = syrup
	return cb
}

// 커피 온도 설정
func (cb *CoffeeBuilder) SetTemperature(temperature string) *CoffeeBuilder {
	cb.coffee.Temperature = temperature
	return cb
}

// 최종적으로 구성된 Coffee 객체 반환
func (cb *CoffeeBuilder) Build() *Coffee {
	return cb.coffee
}
