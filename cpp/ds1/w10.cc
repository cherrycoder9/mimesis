// 다중 상속은 하나의 클래스가 둘 이상의 부모 클래스로부터 멤버(데이터와 함수)를
// 상속받는 메커니즘을 의미함. 코드 재사용성을 극대화하고 복잡한 클래스 계층
// 구조를 설계할 수 있지만 모호성 문제가 같은 복잡성을 동반하기도 함.

#include <iostream>
#include <string>

using namespace std;

class Person {
 public:
  void introduce() const {
    cout << "저는 사람입니다.\n";
  }
};

class Learner {
 public:
  void study() const {
    cout << "공부를 합니다\n";
  }
};

class Student : public Person, public Learner {
 public:
  void describe() const {
    cout << "저는 공부하는 학생입니다.\n";
  }
};

int main(int argc, char const *argv[]) {
  Student student;
  student.introduce();
  student.study();
  student.describe();
  return 0;
}
