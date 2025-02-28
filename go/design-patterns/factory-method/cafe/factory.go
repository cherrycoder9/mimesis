package cafe

// DrinkFactory 인터페이스는 팩토리 메서드 패턴에서 음료 객체를 생성하는 팩토리의 기본 형태 정의
type DrinkFactory interface {
	// CreateDrink 메서드는 Drink 인터페이스를 구현하는 객체를 생성해 반환함
	CreateDrink() Drink
}

// CoffeeFactory는 커피 음료 객체를 생성하는 팩토리
type CoffeeFactory struct{}

// CreateDrink 메서드를 통해 커피 객체를 생성해 반환함
func (cf *CoffeeFactory) CreateDrink() Drink {
	return &Coffee{}
}

// TeaFactory는 차 음료 객체를 생성하는 팩토리
type TeaFactory struct{}

// CreateDrink 메서드를 통해 차 객체 생성해 반환
func (tf *TeaFactory) CreateDrink() Drink {
	return &Tea{}
}
