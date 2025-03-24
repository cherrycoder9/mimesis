package subsystem

import "fmt"

// 설탕을 공급하는 기능을 담당하는 구조체
type SugarDispenser struct{}

// 새로운 SugarDispenser 객체를 생성해 반환
func NewSugarDispenser() *SugarDispenser {
	return &SugarDispenser{}
}

// 설탕을 추가하는 기능
func (sd *SugarDispenser) Dispense() {
	fmt.Println("설탕 디스펜서, 설탕을 추가 중...")
}
