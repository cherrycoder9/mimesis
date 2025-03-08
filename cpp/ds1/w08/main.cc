#include <iostream>

#include "student.hpp"

int main() {
  // student 객체 생성
  Student student("박보람", 12345);

  // 이름과 학번 출력
  cout << "학생 이름: " << student.getName() << "\n";
  cout << "학생 학번: " << student.getId() << "\n";

  // 이름 변경 후 출력
  student.setName("정희철");
  cout << "변경된 이름: " << student.getName() << "\n";

  return 0;
}