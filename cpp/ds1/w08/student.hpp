#ifndef STUDENT_HPP  // 중복 포함 방지
#define STUDENT_HPP

#include <string>
using namespace std;

class Student {
 private:
  string name_;  // 이름
  int id_;       // 학번

 public:
  // 생성자 선언
  Student(const string& name, int id);

  // 멤버 함수 선언
  string getName() const;
  int getId() const;
  void setName(const string& newName);

  // # 소멸자가 필요한 경우는 클래스가 관리해야 할 자원이 있을 경우
  // 동적 할당된 메모리, 파일 핸들, 네트워크 연결같은 외부자원, 스마트포인터
  // 외의 자원 관리
};

#endif  // STUDENT_HPP