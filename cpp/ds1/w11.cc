// 다중 상속의 대표적 문제는 다이아몬드 문제
// 공통 조상 클래스가 여러 경로를 통해 자식 클래스에 상속될때 발생
// 조상 클래스 멤버에 접근할 때 어느 경로를 통해 접근해야 하는지
// 불분명해지는 상황

#include <iostream>

using namespace std;

class Base {
 public:
  void display() const {
    cout << "Base 클래스의 display 함수입니다.\n";
  }
};

class Intermediate1 : public Base {};
class Intermediate2 : public Base {};

class Derived : public Intermediate1, public Intermediate2 {};

int main() {
  Derived derived;
  // derived.display(); // 오류, Derived::display가 모호합니다.
  derived.Intermediate1::display();  // 경시적 경로 지정으로 해결가능

  return 0;
}