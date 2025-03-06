#include <chrono>   // C++11에 추가된 시간 측정 라이브러리
#include <iomanip>  // 소수점 자리 수 조정을 위해 추가
#include <iostream>
#include <string>

using namespace std;

// 값 전달로 계산하는 함수
int addByValue(int a, int b) {
  return a + b;
}

// 참조 전달로 데이터를 수정하는 함수
void swapByReference(int& a, int& b) {
  int temp = a;
  a = b;
  b = temp;
}

// const 참조 전달로 복사 비용을 줄이고 원본 보호
double calculateAverage(const double& num1, const double& num2) {
  return (num1 + num2) / 2.0;
}

// 객체를 반환하는 함수 (복사 생략 RVO 적용 가능)
// 객체를 반환할 경우 복사 생성자가 호출될 수 있음. 그러나 C++11부터 도입된 이동
// 의미론(move semantics)과 복사 생략(RVO: Return Value Optimization) 덕분에
// 성능 최적화가 가능함
string concatenateStrings(const string& str1, const string& str2) {
  return str1 + str2;
}

// C++17 구조적 바인딩을 활용한 복잡한 반환값 처리
tuple<int, double, string> getMultipleValues() {
  return {42, 3.14, "안녕"};
}

int main() {
  auto start = chrono::high_resolution_clock::now();

  // 값 전달 함수 호출
  int sum = addByValue(5, 3);

  // 참조 전달로 값 교환
  int x = 10, y = 20;
  cout << "교환 전: x = " << x << ", y = " << y << "\n";
  swapByReference(x, y);
  cout << "교환 후: x = " << x << ", y = " << y << "\n";

  // const 참조 전달로 평균 계산
  double avg = calculateAverage(10.5, 20.5);
  cout << "두 수의 평균: " << avg << "\n";

  // 객체 반환
  string result = concatenateStrings("안녕, ", "세계!");
  cout << "문자열 연결 결과: " << result << "\n";

  // 구조적 바인딩으로 복수 반환값 처리
  auto [num, pi, greeting] = getMultipleValues();
  cout << "정수: " << num << ", 실수: " << pi << ", 문자열: " << greeting
       << "\n";

  auto end = chrono::high_resolution_clock::now();
  auto duration = chrono::duration_cast<chrono::microseconds>(end - start);
  double duration_ms = duration.count() / 1000.0;
  cout << fixed << setprecision(3);
  cout << "실행시간: " << duration_ms << " 밀리초\n";

  return 0;
}