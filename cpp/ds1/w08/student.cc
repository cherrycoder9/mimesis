#include "student.hpp"

// student 클래스 생성자 정의
Student::Student(const string& name, int id)
    // _를 붙인 이유는 클래스의 멤버변수임을 나타냄
    : name_(name), id_(id) {  // 초기화 리스트를 사용해 멤버 초기화
  // 생성자 본문은 비워둠 (초기화 리스트로 충분)
}

// 학생 이름을 반환하는 함수
string Student::getName() const {
  return name_;
}

// 학번을 반환하는 함수
int Student::getId() const {
  return id_;
}

// 이름을 변경하는 함수
void Student::setName(const string& newName) {
  name_ = newName;
}