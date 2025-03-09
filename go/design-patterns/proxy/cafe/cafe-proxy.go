package cafe

import (
	"fmt"
	"strings"
)

// CafeBene 객체에 접근할 때 중간에서 추가적 로직을 수행하는 구조체
type Proxy struct {
	cafebene   *CafeBene
	authorized bool
	menuCache  map[string]bool
}

// NewProxy는 프록시 객체를 초기화하는 생성자
func NewProxy(auth bool) *Proxy {
	return &Proxy{
		authorized: auth,
		menuCache: map[string]bool{
			"아메리카노": true,
			"라떼":    true,
			"녹차":    true,
		},
	}
}

// Order는 클라이언트 요청을 실제 CafeBene 객체로 전달하기 전에
// 접근 제어와 캐싱 등의 추가 로직을 처리함
func (p *Proxy) Order(drink string) {
	// 접근 권한 확인
	if !p.authorized {
		fmt.Println("🚫 접근 거부: 주문을 위해 로그인하세요.")
		return
	}

	// 메뉴 유효성 체크 (캐싱 사용 예시)
	drink = strings.TrimSpace(drink)
	if available, ok := p.menuCache[drink]; !ok || !available {
		fmt.Printf("⚠️ [%s] 메뉴는 현재 주문이 불가능합니다.\n", drink)
		return
	}

	// 실제 카페베네 객체 지연 초기화
	if p.cafebene == nil {
		fmt.Println("🕑 실제 카페베네 객체를 초기화 중입니다...")
		p.cafebene = &CafeBene{}
	}

	// 실제 객체로 주문 전달
	p.cafebene.Order(drink)
}
