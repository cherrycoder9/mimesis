#include <iostream>
#include <string>

// # 생성자
// 생성자는 클래스 객체가 생성될 때 호출되는 멤버 함수
// 객체의 초기 상태를 설정하는 역할을 함
// 생성자는 클래스 이름과 동일한 이름을 가짐
// 반환 타입을 지정하지 않음
// 객체가 메모리에 할당된 직후 호출되어 멤버 변수의 초기화나 필요한 자원할당
// 수행 기본생성자: 매개변수가 없는 생성자로, 객체를 기본 상태로 초기화 매개변수
// 생성자: 매개변수를 받아 객체를 특정 값으로 초기화 복사 생성자: 동일한
// 클래스의 기존 객체를 복사해 새 객체 생성 이동 생성자: C++11부터 도입된 기능,
// 기존 객체 자원을 이동시켜 새 객체 생성 객체 생성시 기반 클래스의 생성자가
// 먼저 호출된 후 파생 클래스 생성자 호출됨 클래스 내부에서는 멤버 변수가 선언된
// 순서대로 초기화됨 생성자에서 멤버 변수를 초기화할때, 대입 연산자보다는 초기화
// 리스트를 활용하는걸 추천 # 소멸자 객체가 소멸될 때 호출되는 멤버 함수 객체가
// 사용한 자원을 해제하거나 정리 작업을 수행 소멸자는 클래스 이름 앞에 ~를 붙여
// 정의함 매개변수를 받지 않고 반환 타입도 지정하지 않음 소멸자는 스택에서
// 벗어나거나 동적 할당 해제시 자동 호출됨 파생 클래스의 소멸자가 먼저 호출된 후
// 기반 클래스 소멸자가 호출됨

using namespace std;

class Student {
 private:
  string name;
  int age;
  int* scores;  // 동적 할당된 점수 배열
 public:
  // 기본 생성자
  // 콜론 기호는 생성자의 시그니처와 초기화 리스트를 구분하는 역할을 함
  // 콜론 뒤에는 초기화 리스트가 위치해야 하며 멤버 변수를 직접 초기화할때
  // 사용함 불필요한 기본 초기화 후 {}안에서 중복 초기화하지 않도록 함 콜론을
  // 사용하지 않고 {} 본문에서 대입하는 방식은 상수멤버(const)나 참조멤버(&)는
  // 선언과 동시에 초기화해야 하므로 {}방식으로는 사용불가
  Student() : name("이름 없음"), age(0), scores(nullptr) {
    cout << "\n기본 생성자 호출, 학생 객체 생성됨\n";
  }

  // 매개변수 생성자 (초기화리스트 사용)
  Student(const string& studentName, int studentAge)
      : name(studentName), age(studentAge), scores(new int[3]) {
    scores[0] = 90;
    scores[1] = 85;
    scores[2] = 88;
    cout << "\n매개변수 생성자 호출: " << name << " 생성됨\n";
  }

  // 복사 생성자
  Student(const Student& other)
      : name(other.name), age(other.age), scores(new int[3]) {
    for (size_t i = 0; i < 3; i++) {
      scores[i] = other.scores[i];  // 깊은 복사 수행
    }
    cout << "\n복사 생성자 호출: " << name << " 복사됨\n";
  }

  // 이동 생성자 (C++ 이상)
  // noexcept, 함수가 예외를 던지지 않음을 보장하기 위해 사용하는 키워드
  // 함수 선언 뒤에 붙일수 있음, 특정 조건에 따라 동작하게 할수도 있음
  // noexcept(someCondition)
  Student(Student&& other) noexcept
      // string은 동적으로 메모리를 관리하는 클래스 타입임
      // 내부적으로 힙에 문자열 데이터를 저장하고 포인터로 가리킴
      // 이동 생성자의 목적은 자원의 소유권을 이전하는 것임
      // string의 경우 move를 사용하면 기존 객체가 소유한 메모리 포인터를
      // 새 객체로 옮기고 원본 객체의 포인터를 nullptr로 설정해
      // 자원이 중복 해제되는걸 방지
      : name(move(other.name)), age(other.age), scores(other.scores) {
    other.scores = nullptr;
    cout << "\n이동 생성자 호출: " << name << "이동됨\n";
  }

  // 소멸자
  ~Student() {
    delete[] scores;
    cout << "\n소멸자 호출: " << name << " 소멸됨\n";
  }

  // 학생 정보 출력 함수
  void printInfo() const {
    cout << "이름: " << name << ", 나이: " << age;
    if (scores) {
      cout << ", 점수: " << scores[0] << ", " << scores[1] << ", " << scores[2];
    }
    cout << "\n";
  }
};

int main(int argc, char const* argv[]) {
  // 기본 생성자 테스트
  Student student1;
  student1.printInfo();

  // 매개변수 생성자 테스트
  Student student2("박정호", 20);
  student2.printInfo();

  // 복사 생성자 테스트
  Student student3 = student2;
  student3.printInfo();

  // 이동 생성자 테스트
  Student student4 = move(student2);
  student4.printInfo();
  student2.printInfo();  // 이동후 student2의 상태 확인

  return 0;
}
