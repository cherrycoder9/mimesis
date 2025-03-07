#include <iostream>
#include <string>

using namespace std;

// 함수 오버로딩은 C++에서 동일한 이름을 가진 여러 함수를 정의하는 기능
// 객체지향 프로그래밍과 제네릭 프로그래밍의 기반을 이루는 개념 중 하나
// 오버로딩은 함수의 이름은 동일하되, 매개변수의 개수나 타입, 혹은 둘 다를
// 다르게 함으로써 컴파일러가 함수 호출시 적절한 버전을 선택할 수 있도록 함
// C++에서 함수 오버로딩이 가능하려면 컴파일러가 함수 호출 시각에 각 호출을
// 고유하게 식별할 수 있어야 하며 이를 함수 시그니처라고 부름
// 매개변수의 개수 또는 타입이 달라야 함
// 반환 타입만 다른 경우 오버로딩으로 간주되지 않음

void print(int number) {
  cout << "정수: " << number << endl;
}

void print(double number) {
  cout << "실수: " << number << endl;
}

void print(const string& text) {
  // const 참조를 사용해 불필요한 복사 방지
  cout << "문자열: " << text << endl;
}

int main(int argc, char const* argv[]) {
  int intValue = 42;
  double doubleValue = 3.14;
  string stringValue = "헬로 C++";

  // 오버로드된 함수 호출
  print(intValue);
  print(doubleValue);
  print(stringValue);

  return 0;
}
