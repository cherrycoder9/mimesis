// C++에서 가상 함수는 가상 함수 테이블(VTable)을 통해 구현됨
// VTable은 각 클래스마다 생성되는 함수 포인터 배열
// 해당 클래스의 가상 함수들에 대한 주소를 저장함
// 객체가 생성될때 객체 내부에는 VTable을 가리키는 포인터(VPtr)가 포함됨
// 가상 함수 호출시 런타임에 VPtr을 통해 VTable을 참조하고
// 실제 호출할 함수의 주소를 찾아 실행함
// 이러한 매커니즘은 동적 디스패치라고도 불리며 컴파일 타임이 아닌
// 런타임에 함수 호출을 결정함. 약간의 성능 오버헤드를 초래하지만
// 다형성을 통한 코드 유연성과 확장성이 있게됨

#include <iostream>
#include <memory>
#include <string>

using namespace std;

class Animal {
 public:
  // 가상 함수 speak 선언
  virtual void speak() const {
    cout << "동물이 소리를 냅니다." << "\n";
  }

  // 가상 소멸자 선언
  // 파생 클래스의 소멸자가 호출되도록 보장함
  virtual ~Animal() = default;
};

// 파생 클래스 Dog 정의
class Dog : public Animal {
 public:
  // speak 함수 재정의
  // override, 가상 함수를 재정의함을 명시 (C++11 이상)
  void speak() const override {
    cout << "멍멍!" << "\n";
  }
};

// 파생 클래스 Cat 정의
class Cat : public Animal {
 public:
  // speak 함수 재정의
  void speak() const override {
    cout << "야옹!" << "\n";
  }
};

int main(int argc, char const *argv[]) {
  // 스마트 포인터를 사용해 메모리 누수 방지
  unique_ptr<Animal> animal1 = make_unique<Dog>();
  unique_ptr<Animal> animal2 = make_unique<Cat>();

  // 가상 함수 호출
  animal1->speak();
  animal2->speak();

  // unique_ptr은 자동으로 메모리 해제
  return 0;
}
