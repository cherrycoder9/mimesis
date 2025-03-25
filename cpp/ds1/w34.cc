// throw는 예외를 명시적으로 발생시키는데 사용되는 키워드
// C++에서는 예외로 던질 수 있는 값의 타입에 제약이 없지만
// 표준 라이브러리의 std::exception 클래스를 상속받은 객체를
// 사용하는 것이 일반적임
// throw가 호출되면 프로그램의 제어 흐름은 즉시 해당 예외를
// 처리할 수 있는 catch 블록으로 이동함
#include <cstdlib>
#include <iostream>
#include <stdexcept>

using namespace std;

int main(int argc, char const* argv[]) {
  try {
    // 예외를 발생시키는 상황
    throw runtime_error("런타임 오류 발생");
  } catch (const exception& e) {
    // 발생한 예외를 처리
    cerr << "예외발생: " << e.what() << '\n';
  }

  return EXIT_SUCCESS;
}
