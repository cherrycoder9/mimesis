// 함수 템플릿은 C++에서 제네릭 프로그래밍을 구현하는 핵심 도구
// 특정 데이터 타입에 의존하지 않고 일반화된 함수를 작성할 수 있게함
// 함수 템플릿은 컴파일 시점에 필요한 타입으로 구체화 됨
// template 키워드로 정의되며 타입 매개변수를 사용해 타입을 추상화
// 컴파일러는 함수 호출 시점에 전달된 인자의 타입을 바탕으로 적절한 함수 생성

// 템플릿은 모든 타입에 대해 동작한다고 보장할 수 없음
// 템플릿 내부에서 사용하는 연산이 특정 타입에서 지원되지 않을 경우
// 컴파일 오류가 발생할 수 있음. 이것을 해결하기 위해 C++20 부터는 concept를
// 활용해 타입 제약을 명시할 수 있음

#include <iostream>
#include <string>

using namespace std;

// 두 값을 비교해 더 큰 값을 반환하는 함수 템플릿
template <typename T>
T getMax(T a, T b) {
  return (a > b) ? a : b;
}

int main(int argc, char const *argv[]) {
  int intResult = getMax(5, 3);
  cout << "두 정수 중 큰값: " << intResult << "\n";

  double doubleResult = getMax(4.7, 9.2);
  cout << "두 실수 중 큰값: " << doubleResult << "\n";

  string strResult = getMax(string("사과"), string("바나나"));
  cout << "두 문자열 중 사전순으로 뒤에 오는 값: " << strResult << "\n";

  return 0;
}
