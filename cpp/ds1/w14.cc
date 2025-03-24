// 추상 클래스는 하나 이상의 순수 가상 함수를 포함하는 클래스
// 직접적으로 객체를 생성할 수 없는 클래스임
// 구체적인 구현을 제공하기보단 공통된 인터페이스나 설계 템플릿을 정의하는 역할
// 추상클래스는 상속을 통해 파생클래스에서 구체적 동작을 구현하도록 강제함
// 다형성은 다양한 객체가 동일한 인터페이스를 공유하면서도 각기
// 다른 방식으로 동작할 수 있게 해줌

// # 추상 클래스 특징
// 객체 생성 불가, 추상클래스는 미완성된 설계도로 간주되며 직접 인스턴화 불가
// 순수 가상 함수 포함, 순수 가상 함수는 함수 선언 뒤에 = 0을 붙여 정의함
// 몸체가 없는 함수이며 추상 클래스는 최소 하나 이상의 순수가상함수를 가져야함
// 상속 강제성, 추상클래스는 상속받은 파생 클래스는 모든 순수가상함수를
// 반드시 구현해야 함. 그렇지 않으면 파생 클래스 역시 추상클래스가 되어 객체생성
// 불가능
// 다형성 지원, 추상 클래스의 포인터나 참조를 사용해 파생 클래스의 객체를
// 조작할 수 있어 런타임에 동적으로 적절한 동작을 호출할 수 있음

#include <iostream>
#include <memory>
#include <string>

using namespace std;

class Animal {
 public:
  // 순수 가상 함수, 모든 파생 클래스가 반드시 구현해야 함
  virtual void makeSound() const = 0;

  // 가상 소멸자, 다형적 소멸을 보장해 메모리 누수 방지
  virtual ~Animal() = default;
};

class Dog : public Animal {
 public:
  // 순수 가상 함수 오버라이드
  void makeSound() const override {
    cout << "멍멍!" << "\n";
  }
};

class Cat : public Animal {
 public:
  void makeSound() const override {
    cout << "야옹!" << "\n";
  }
};

int main(int argc, char const *argv[]) {
  // 스마트 포인터를 사용해 동적 할당 및 자동 해제
  unique_ptr<Animal> dog = make_unique<Dog>();
  unique_ptr<Animal> cat = make_unique<Cat>();

  // 다형성을 활용해 각 객체의 소리 출력
  dog->makeSound();
  cat->makeSound();

  return 0;
}
