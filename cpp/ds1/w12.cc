// 가상 상속으로 다이아몬드 문제 해결
// 가상 상속을 사용하면 공통 조상 클래스의 인스턴스가 단일하게 유지되어 모호성이
// 제거됨 중간 클래스에서 virtual 키워드를 사용해 상속을 선언함으로써 구현됨

#include <iostream>

using namespace std;

class Base {
 public:
  void display() const {
    cout << "Base 클래스의 display 함수입니다.\n";
  }
};

// 가상 상속은 다중 상속의 설계를 단순화하지만, 객체 생성시 약간의 성능
// 오버헤드가 발생할 수 있으므로 필요에 따라 신중히 사용
class Intermediate1 : virtual public Base {
  // Base를 가상 상속
};

class Intermediate2 : virtual public Base {
  // Base를 가상 상속
};

class Derived : public Intermediate1, public Intermediate2 {};

int main(int argc, char const *argv[]) {
  Derived derived;
  derived.display();  // 모호성 없이 정상 호출
  return 0;
}
