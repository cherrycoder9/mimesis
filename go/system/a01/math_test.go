package math

import "testing"

// go test
// Go 코드에 대한 테스트를 자동으로 탐색하고 실행해주는 명령어
// 테스트 파일(*_test.go), 테스트 함수(TestXxx)를 자동으로 찾아 실행
func TestAdd(t *testing.T) {
	result := Add(2, 2)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
