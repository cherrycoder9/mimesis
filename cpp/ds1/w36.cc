// 예외 발생시 현재 try 블록에 적합한 catch 블록이 없으면 예외는
// 호출 스택을 따라 상위 함수로 전파됨
// 예외를 발생시킨 함수와 다른 위치에서 처리할 수 있음
#include <cstdlib>
#include <iostream>
#include <stdexcept>

using namespace std;

void innerFunction() {
  // 내부 함수에서 예외 발생
  throw logic_error("논리 오류 발생");
}

void outerFunction() {
  try {
    innerFunction();
  } catch (const logic_error& e) {
    // 상위 함수에서 예외 처리
    cerr << "논리 오류: " << e.what() << '\n';
  }
}

int main(int argc, char const* argv[]) {
  outerFunction();

  return EXIT_SUCCESS;
}
