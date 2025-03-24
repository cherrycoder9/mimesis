// # 템플릿 특수화
// 제네릭 프로그래밍을 할때 템플릿의 일반화된 동작을
// 특정 타입이나 조건에 맞춰 세부적으로 조정하는 기법
// 완전 특수화와 부분 특수화로 크게 두가지가 있음

// * 완전 특수화
// 템플릿의 모든 매개변수를 특정 타입으로 고정해 정의하는 방식
// 일반 템플릿 대신 특정 타입에 대해 완전히 다른 구현을 하고자 할때

// * 부분 특수화
// 템플릿 매개변수 중 일부만 특정 조건에 맞춰 정의하는 방식
// 클래스 템플릿에서 주로 사용됨 (함수 템플릿은 부분 특수화 미지원)
// 특정 패턴이나 제약을 가진 타입에 대해 세부적 동작을 지정할 때 적합

// 오버로딩 vs 특수화: 함수 템플릿의 경우 부분 특수화를 사용할 수 없으므로
// 오버로딩으로 대체해야 하는 경우가 많음
// 컴파일러 선택: 컴파일러는 가장 구체적인 특수화된 버전을 우선적으로 선택함.
// Container<int>가 int에 대해 호출되면 일반 템플릿 대신 특수화된 버전이 실행됨
#include <iostream>
#include <string>

using namespace std;

// 기본 템플릿 클래스 정의
template <typename T>
class Container {
 public:
  void print(T value) {
    cout << "기본 템플릿: " << value << "\n";
  }
};

// 완전 특수화, int 타입에 대한 특수화
template <>
class Container<int> {
 public:
  void print(int value) {
    cout << "int 특수화: " << value << " (정수형)" << "\n";
  }
};

// 부분 특수화, 포인터 타입(T*)에 대한 특수화
template <typename T>
class Container<T*> {
 public:
  void print(T* value) {
    cout << "포인터 특수화: " << *value << " (포인터 값)" << "\n";
  }
};

int main(int argc, char const* argv[]) {
  // 기본 템플릿 사용
  Container<string> stringContainer;
  stringContainer.print("안녕하세요");

  // 완전 특수화 사용
  Container<int> intContainer;
  intContainer.print(42);

  // 부분 특수화 사용
  int number = 100;
  Container<int*> pointerContainer;
  pointerContainer.print(&number);

  return 0;
}
