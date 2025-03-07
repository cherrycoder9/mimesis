#include <algorithm>
#include <iostream>
#include <vector>

// 람다 표현식은 C++11에서 도입된 기능
// 이름 없는 함수(익명 함수)를 정의할 수 있게 함
// 짧고 간단한 작업을 정의하거나, 함수 객체를 즉석에서 생성해야 할 때 사용

// 람다 표현식은 내부적으로 컴파일러에 의해 함수 객체(즉, operator()를
// 오버로드한 클래스)로 변환됨. 이것을 클로저라고 부르며, 캡처된 변수는 이
// 클로저 객체의 멤버로 저장됨

// # 캡처 방식
// 람다 표현식의 캡처 절은 외부 변수 접근 방식을 정의함
// 값 캡처 [x]: 변수 x를 값으로 캡쳐, 기본적으로 상수로 취급
// 참조 캡처 [&x]: 변수 x를 참조로 캡처해 수정 가능함
// 전체 값 캡처 [=]: 람다 외부의 모든 변수를 값으로 캡처함
// 전체 참조 캡처 [&]: 람다 외부의 모든 변수를 참조로 캡처함
// 혼합 캡처: [=,&x]는 x를 참조로, 나머지를 값으로 캡처함

// 주로 다음과 같은 상황에서 활용함
// 1. STL 알고리즘에 사용자 정의 동작을 전달할 때
// 2. 비동기 작업이나 스레드에서 간단한 작업 정의할 때
// 3. 이벤트 핸들러나 콜백 함수를 즉석에서 작성할 때

// # 주의점
// 캡처된 참조는 람다 호출시 외부 변수가 유효해야 함 (댕글링 참조 주의)
// 과도한 캡처는 코드 복잡성을 증가시키므로 필요한 변수만 캡처 권장

using namespace std;

int main(int argc, char const *argv[]) {
  // 숫자 벡터 생성
  vector<int> numbers{5, 2, 9, 1, 5, 6};

  // 1. 기본 람다: 벡터를 오름차순 정렬
  sort(numbers.begin(), numbers.end(), [](int a, int b) { return a < b; });

  // 정렬된 결과 출력
  cout << "오름차순 결과: ";
  for (int num : numbers) {
    cout << num << " ";
  }
  cout << "\n";

  // 2. 캡처 사용, 특정 값보다 큰 숫자만 출력
  int threshold{4};
  cout << threshold << "보다 큰 숫자: ";
  for_each(numbers.begin(), numbers.end(), [threshold](int num) {
    if (num > threshold) {
      cout << num << " ";
    }
  });
  cout << "\n";

  // 3. 참조 캡처, 외부 변수 수정
  int sum{0};
  for_each(numbers.begin(), numbers.end(), [&sum](int num) {  // 참조로 캡처
    sum += num;
  });
  cout << "숫자의 합: " << sum << "\n";

  // 4. 반환 타입 명시, 두수의 최대값 계산
  auto max_lambda = [](int a, int b) -> int {  // 반환 타입 명시
    return (a > b) ? a : b;
  };
  cout << "5와 7 중 큰 값: " << max_lambda(5, 7) << "\n";

  // C++17 이후로 람다에 constexpr 붙여 컴파일 타임 계산 가능해짐
  return 0;
}
