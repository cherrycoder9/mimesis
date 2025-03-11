package subsystem

import "fmt"

// 우유를 데우고 거품을 만드는 기능을 담당하는 구조체
type MilkSteamer struct{}

// 새로운 MilkSteamer 객체를 생성해 반환
func NewMilkSteamer() *MilkSteamer {
	return &MilkSteamer{}
}

// 우유를 데우고 거품을 만드는 기능
func (ms *MilkSteamer) Steam() {
	fmt.Println("우유 스티머, 우유를 스팀 중...")
}
