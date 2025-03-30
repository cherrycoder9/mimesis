// 사용자 정의 예외 클래스
#include <cstdlib>
#include <iostream>
#include <stdexcept>
#include <string>

using namespace std;

// 사용자 정의 예외 클래스 정의
class DivisionByZeroException : public exception {
 private:
  string message;  // 예외 메시지를 저장하는 멤버 변수

 public:
  // 생성자, 예외 메시지를 인자로 받아 초기화
  explicit DivisionByZeroException(const string& msg) : message(msg) {
  }

  // what() 메서드 오버라이딩, 예외 메시지 반환
  const char* what() const noexcept override {
    return message.c_str();
  }
};

// 0으로 나누기를 시도하는 함수
double divide(double numerator, double denominator) {
  if (denominator == 0) {
    throw DivisionByZeroException("0으로 나눌 수 없습니다.");
  }
  return numerator / denominator;
}

int main(int argc, char const* argv[]) {
  try {
    // 정상적인 나눗셈 테스트
    double result = divide(10.0, 2.0);
    cout << "결과: " << result << endl;

    // 예외를 유발하는 나눗셈 테스트
    result = divide(5.0, 0.0);
  } catch (const DivisionByZeroException& e) {
    cerr << "오류발생: " << e.what() << '\n';
    return EXIT_FAILURE;
  } catch (const exception& e) {
    cerr << "알수 없는 오류: " << e.what() << '\n';
    return EXIT_FAILURE;
  }

  return EXIT_SUCCESS;
}
