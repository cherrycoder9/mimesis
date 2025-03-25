// noexcept 키워드
// C++11부터 도입된 noexcept는 함수가 예외를 발생시키지 않음을 명시
// 컴파일러가 코드를 최적화하는데 도움을 주며 예외 안전성을 보장함

#include <cstdlib>
#include <iostream>

using namespace std;

void noExceptionFunction() noexcept {
  cout << "Exception 없는 함수" << endl;
  // throw runtime_error("에러");
}

int main(int argc, char const* argv[]) {
  try {
    noExceptionFunction();
  } catch (const exception& e) {
    cerr << e.what() << '\n';
  }

  cout << "프로그램 정상 종료" << endl;

  return EXIT_SUCCESS;
}
