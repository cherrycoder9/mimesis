// try 블록은 예외가 발생할 가능성이 있는 코드를 감싸는 역할
// try 블록 내에서 예외 발생시 프로그램은 즉시 해당 예외를
// 처리할 수 있는 catch 블록으로 제어를 넘김
// try 블록은 반드시 하나 이상의 catch 블록과 함께 사용됨
#include <cstdlib>
#include <iostream>
#include <stdexcept>

using namespace std;

int main(int argc, char const* argv[]) {
  try {
    int denominator = 0;
    if (denominator == 0) {
      throw invalid_argument("0으로 나눌 수 없습니다.");
    }
  } catch (const invalid_argument& e) {
    // 특정 예외 타입 처리
    cerr << "잘못된 인자: " << e.what() << '\n';
  } catch (const exception& e) {
    // 기타 예외 처리
    cerr << "기타 예외: " << e.what() << endl;
  }

  return EXIT_SUCCESS;
}
