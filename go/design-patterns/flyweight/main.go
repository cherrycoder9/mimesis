package main

import (
	"fmt"
)

// # 플라이웨이트 패턴
// 동일하거나 유사한 객체를 공유함으로써 메모리 사용을 줄이는 방식
// 많은 수의 작은 객체를 생성해야 할 때 효과적임
// 게임에서 반복되는 캐릭터 텍스처나 폰트 글리프를 공유할 수 있음
// 불변 데이터를 캐싱하고 가변 데이터는 외부에서 관리함
// 객체 상태가 자주 바뀌면 공유가 어려워질수 있음
func main() {
	// Go 1.20 부터는 전역 난수 생성기가 이미 암호화된 안전한 값으로 자동 초기화되므로
	// 단순히 난수를 얻기 위해 rand.Seed를 호출할 필요가 없어짐
	// rand.Seed(time.Now().UnixNano())

	factory := GetFactoryInstance()

	orders := []struct {
		beverageName string
		size         string
	}{
		{"아메리카노", "톨"},
		{"라떼", "그란데"},
		{"프라푸치노", "벤티"},
		{"아메리카노", "벤티"},
		{"라떼", "톨"},
		{"아메리카노", "그란데"},
	}

	fmt.Println("카페베네 주문 내역 처리 시작")
	for idx, order := range orders {
		fmt.Printf("\n[주문 %d] '%s(%s)' 주문이 들어왔습니다.\n",
			idx+1, order.beverageName, order.size)
		bev := factory.GetBeverage(order.beverageName)
		bev.Serve(order.size)
	}

	fmt.Printf("\n✅ 모든 주문 처리 완료!\n")
}
