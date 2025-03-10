package main

import (
	"fmt"
	"sync"
)

// Flyweight 객체를 관리하고 재사용함
type BeverageFactory struct {
	beverages map[string]Beverage // Flyweight 객체를 저장할 맵
	mutex     sync.Mutex          // 고루틴 동시성 보호를 위한 mutex
}

// 전역 팩토리 인스턴스 생성 (싱글톤)
var factoryInstance *BeverageFactory
var once sync.Once

// BeverageFactory의 싱글톤 인스턴스를 반환함
func GetFactoryInstance() *BeverageFactory {
	once.Do(func() {
		factoryInstance = &BeverageFactory{
			beverages: make(map[string]Beverage),
		}
	})
	return factoryInstance
}

// 음료 이름을 받아 Flyweight 객체를 반환함
// 이미 존재하면 기존객체 반환, 없으면 생성해 맵에 저장
func (f *BeverageFactory) GetBeverage(name string) Beverage {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if bev, exists := f.beverages[name]; exists {
		fmt.Printf("[공유 객체 재사용] 기존의 '%s' 음료 객체 사용\n", name)
		return bev
	}

	var description string
	switch name {
	case "아메리카노":
		description = "진한 에스프레소에 물을 더한 음료"
	case "라떼":
		description = "부드러운 우유와 에스프레소의 조합"
	case "프라푸치노":
		description = "달콤하고 차가운 블렌딩 음료"
	default:
		description = "맛있는 카페베네 음료"
	}

	newBeverage := &ConcreteBeverage{name: name, description: description}
	f.beverages[name] = newBeverage
	fmt.Printf("[새로운 공유 객체 생성] '%s' 음료 객체 생성\n", name)

	return newBeverage
}
